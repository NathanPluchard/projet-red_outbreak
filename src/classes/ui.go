package classes

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var ansiPattern = regexp.MustCompile("\x1b\\[[0-9;]*m")

// stripANSI retire les codes couleur/style avant de mesurer une chaîne.
func stripANSI(s string) string {
	return ansiPattern.ReplaceAllString(s, "")
}

// runeWidth estime la largeur d'affichage d'un caractère dans un
// terminal. Les émojis et symboles larges comptent pour 2 colonnes,
// les sélecteurs de variante/ZWJ pour 0, le reste pour 1.
func runeWidth(r rune) int {

	switch {

	case r == 0xFE0F || r == 0x200D:
		return 0

	case r >= 0x1F000 && r <= 0x1FFFF:
		return 2

	case r >= 0x2600 && r <= 0x27BF:
		return 2

	case r >= 0x2B00 && r <= 0x2BFF:
		return 2
	}

	return 1

}

// visibleWidth calcule la largeur réellement affichée d'une chaîne
// (codes ANSI ignorés, caractères larges comptés double). Utilisé
// pour que les cadres (Panel, TitleBox) restent bien alignés même
// avec des couleurs et des emojis dedans.
func visibleWidth(s string) int {

	width := 0

	for _, r := range stripANSI(s) {
		width += runeWidth(r)
	}

	return width

}

const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Dim    = "\033[2m"
	Italic = "\033[3m"

	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"
)

func ClearScreen() {
	fmt.Print("\033[2J\033[H")
}

func Pause() {
	fmt.Print(
		"\n" +
			Dim +
			"Appuyez sur Entrée pour continuer..." +
			Reset,
	)

	Reader.ReadString('\n')

}

func ratio(current, maximum int) float64 {

	if maximum <= 0 {
		return 0
	}

	r := float64(current) / float64(maximum)

	if r < 0 {
		return 0
	}

	if r > 1 {
		return 1
	}

	return r

}

func bar(
	current int,
	maximum int,
	width int,
	color string,
) string {

	r := ratio(current, maximum)

	filled := int(r * float64(width))

	if filled > width {
		filled = width
	}

	if filled < 0 {
		filled = 0
	}

	return color + strings.Repeat("█", filled) + Reset + Dim + strings.Repeat("░", width-filled) + Reset

}

func HPBar(current, maximum int) string {

	r := ratio(current, maximum)

	color := Green

	if r <= 0.25 {
		color = BrightRed
	} else if r <= 0.5 {
		color = Yellow
	}

	return "[" +
		bar(
			current,
			maximum,
			20,
			color,
		) +
		"]"

}

func ManaBar(current, maximum int) string {

	return "[" +
		bar(
			current,
			maximum,
			20,
			Cyan,
		) +
		"]"

}

func ExpBar(current, maximum int) string {

	return "[" +
		bar(
			current,
			maximum,
			20,
			Magenta,
		) +
		"]"

}

func Separator() string {

	return Dim +
		strings.Repeat("─", 58) +
		Reset

}

func TitleBox(title string) string {

	inner := " " + title + " "

	width := visibleWidth(inner)

	top :=
		"╔" +
			strings.Repeat("═", width) +
			"╗"

	middle :=
		"║" +
			inner +
			"║"

	bottom :=
		"╚" +
			strings.Repeat("═", width) +
			"╝"

	return Bold + BrightCyan + top + "\n" + middle + "\n" + bottom + Reset

}

func Panel(
	title string,
	lines ...string,
) string {

	var builder strings.Builder

	const innerWidth = 56

	titleWidth := visibleWidth(title)

	dashes := innerWidth - 1 - titleWidth

	if dashes < 1 {
		dashes = 1
	}

	builder.WriteString(Bold)
	builder.WriteString(BrightBlue)
	builder.WriteString("┌─ ")
	builder.WriteString(title)
	builder.WriteString(" ")
	builder.WriteString(strings.Repeat("─", dashes))
	builder.WriteString("┐")
	builder.WriteString(Reset)
	builder.WriteString("\n")

	for _, line := range lines {

		lineWidth := visibleWidth(line)

		padding := innerWidth - lineWidth

		if padding < 0 {
			padding = 0
		}

		builder.WriteString("│ ")
		builder.WriteString(line)
		builder.WriteString(strings.Repeat(" ", padding))
		builder.WriteString(" │\n")
	}

	builder.WriteString(Bold)
	builder.WriteString(BrightBlue)
	builder.WriteString("└")
	builder.WriteString(strings.Repeat("─", innerWidth+2))
	builder.WriteString("┘")
	builder.WriteString(Reset)

	return builder.String()

}

// StatLine formate une ligne "label : valeur" avec le label aligné,
// prête à être passée à Panel(). labelWidth est la largeur (en
// colonnes visibles) réservée au label avant le ":".
func StatLine(label string, value string) string {

	const labelWidth = 10

	pad := labelWidth - visibleWidth(label)

	if pad < 0 {
		pad = 0
	}

	return Bold +
		label +
		strings.Repeat(" ", pad) +
		Reset +
		" : " +
		value

}

func Colorize(
	color string,
	text string,
) string {

	return color +
		text +
		Reset
}

func TypeWriter(text string, delay time.Duration) {
	for _, c := range text {
		fmt.Print(string(c))
		time.Sleep(delay)
	}
	fmt.Println()
}