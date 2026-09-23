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

// RGB renvoie un code couleur ANSI 24 bits (vraie couleur), pour
// des teintes précises et des dégradés que la palette 16 couleurs
// ne permet pas.
func RGB(r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

var (
	ThemeNeonGreen   = RGB(57, 255, 90)
	ThemeToxicYellow = RGB(212, 255, 0)
	ThemeBloodRed    = RGB(255, 40, 60)
	ThemeRustOrange  = RGB(255, 140, 40)
	ThemeInfected    = RGB(175, 70, 255)
	ThemeSkyCyan     = RGB(80, 210, 255)
	ThemeGold        = RGB(255, 195, 40)
	ThemeBone        = RGB(230, 230, 220)
)

func lerp(a, b int, t float64) int {
	return int(float64(a) + float64(b-a)*t)
}

// gradientColor interpole entre deux couleurs RGB selon t (0 à 1).
func gradientColor(from, to [3]int, t float64) string {

	if t < 0 {
		t = 0
	}

	if t > 1 {
		t = 1
	}

	return RGB(
		lerp(from[0], to[0], t),
		lerp(from[1], to[1], t),
		lerp(from[2], to[2], t),
	)

}

func maxInt(a, b int) int {

	if a > b {
		return a
	}

	return b

}

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

var barBlocks = []rune{' ', '▏', '▎', '▍', '▌', '▋', '▊', '▉'}

// gradientBar dessine une barre de progression en dégradé de
// couleur (from → to selon le taux de remplissage), avec une
// précision au demi-caractère près pour un rendu plus fin qu'un
// simple blocage par caractère entier.
func gradientBar(
	current int,
	maximum int,
	width int,
	from [3]int,
	to [3]int,
) string {

	r := ratio(current, maximum)

	exact := r * float64(width)

	full := int(exact)

	if full > width {
		full = width
	}

	frac := exact - float64(full)

	var b strings.Builder

	b.WriteString(gradientColor(from, to, r))
	b.WriteString(strings.Repeat("█", full))

	if full < width && frac > 0 {

		idx := int(frac * float64(len(barBlocks)))

		if idx >= len(barBlocks) {
			idx = len(barBlocks) - 1
		}

		b.WriteRune(barBlocks[idx])

		full++
	}

	b.WriteString(Reset)

	if width-full > 0 {
		b.WriteString(Dim)
		b.WriteString(strings.Repeat("░", width-full))
		b.WriteString(Reset)
	}

	return b.String()

}

func HPBar(current, maximum int) string {

	return "[" +
		gradientBar(
			current,
			maximum,
			20,
			[3]int{235, 45, 65},
			[3]int{60, 225, 95},
		) +
		"]"

}

func ManaBar(current, maximum int) string {

	return "[" +
		gradientBar(
			current,
			maximum,
			20,
			[3]int{40, 80, 200},
			[3]int{90, 220, 255},
		) +
		"]"

}

func ExpBar(current, maximum int) string {

	return "[" +
		gradientBar(
			current,
			maximum,
			20,
			[3]int{130, 40, 200},
			[3]int{225, 120, 255},
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

	from := [3]int{57, 255, 90}
	to := [3]int{255, 40, 60}

	var top strings.Builder
	var bottom strings.Builder

	top.WriteString("╔")
	bottom.WriteString("╚")

	for i := 0; i < width; i++ {

		t := float64(i) / float64(maxInt(width-1, 1))

		top.WriteString(gradientColor(from, to, t))
		top.WriteString("═")

		bottom.WriteString(gradientColor(to, from, t))
		bottom.WriteString("═")
	}

	top.WriteString(Reset)
	top.WriteString(gradientColor(from, to, 1))
	top.WriteString("╗")
	top.WriteString(Reset)

	bottom.WriteString(Reset)
	bottom.WriteString(gradientColor(to, from, 1))
	bottom.WriteString("╝")
	bottom.WriteString(Reset)

	middle :=
		ThemeBone + "║" + Reset +
			Bold + BrightWhite + inner + Reset +
			ThemeBone + "║" + Reset

	return top.String() + "\n" + middle + "\n" + bottom.String()

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
	builder.WriteString(ThemeSkyCyan)
	builder.WriteString("┌─ ")
	builder.WriteString(Reset)
	builder.WriteString(Bold)
	builder.WriteString(ThemeGold)
	builder.WriteString(title)
	builder.WriteString(Reset)
	builder.WriteString(Bold)
	builder.WriteString(ThemeSkyCyan)
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

		builder.WriteString(ThemeSkyCyan)
		builder.WriteString("│ ")
		builder.WriteString(Reset)
		builder.WriteString(line)
		builder.WriteString(strings.Repeat(" ", padding))
		builder.WriteString(ThemeSkyCyan)
		builder.WriteString(" │")
		builder.WriteString(Reset)
		builder.WriteString("\n")
	}

	builder.WriteString(Bold)
	builder.WriteString(ThemeSkyCyan)
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

var logoLines = []string{
	`  ___  _   _ _____ ____  ____  _____    _    _  __`,
	` / _ \| | | |_   _| __ )|  _ \| ____|  / \  | |/ /`,
	`| | | | | | | | | |  _ \| |_) |  _|   / _ \ | ' / `,
	`| |_| | |_| | | | | |_) |  _ <| |___ / ___ \| . \ `,
	` \___/ \___/  |_| |____/|_| \_\_____/_/   \_\_|\_\`,
}

// Logo renvoie le grand titre "OUTBREAK" en ASCII art, avec un
// dégradé vertical du vert toxique vers le rouge infecté.
func Logo() string {

	var b strings.Builder

	from := [3]int{57, 255, 90}
	to := [3]int{255, 40, 60}

	for i, line := range logoLines {

		t := float64(i) / float64(maxInt(len(logoLines)-1, 1))

		b.WriteString(gradientColor(from, to, t))
		b.WriteString(Bold)
		b.WriteString(line)
		b.WriteString(Reset)
		b.WriteString("\n")
	}

	return b.String()

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
