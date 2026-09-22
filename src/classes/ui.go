package classes

import (
	"fmt"
	"strings"
<<<<<<< HEAD
=======
	"time"
>>>>>>> origin/main
)

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

func max(a, b int) int {
	if a > b {
		return a
	}

	return b

}

<<<<<<< HEAD
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

=======
func bar(current, max, width int, color string) string {
	r := ratio(current, max)
>>>>>>> origin/main
	filled := int(r * float64(width))

	if filled > width {
		filled = width
	}

	if filled < 0 {
		filled = 0
	}

	return color + strings.Repeat("█", filled) + Reset + Dim + strings.Repeat("░", width-filled) + Reset

}

<<<<<<< HEAD
func HPBar(current, maximum int) string {

	r := ratio(current, maximum)

=======
func HPBar(current, max int) string {
	r := ratio(current, max)
>>>>>>> origin/main
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

<<<<<<< HEAD
func ManaBar(current, maximum int) string {

	return "[" +
		bar(
			current,
			maximum,
			20,
			Cyan,
		) +
		"]"

=======
func ManaBar(current, max int) string {
	return "[" + bar(current, max, 20, Cyan) + "]"
>>>>>>> origin/main
}

func ExpBar(current, maximum int) string {

<<<<<<< HEAD
	return "[" +
		bar(
			current,
			maximum,
			20,
			Magenta,
		) +
		"]"

}

=======
>>>>>>> origin/main
func Separator() string {

<<<<<<< HEAD
	return Dim +
		strings.Repeat("─", 58) +
		Reset

}

=======
>>>>>>> origin/main
func TitleBox(title string) string {

	inner := " " + title + " "

	width := len([]rune(inner))

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

	titleLength := len([]rune(title))

	dashes := 54 - titleLength

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

		lineLength := len([]rune(line))

		padding := 56 - lineLength

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
	builder.WriteString(strings.Repeat("─", 58))
	builder.WriteString("┘")
	builder.WriteString(Reset)

	return builder.String()

}

func Colorize(
	color string,
	text string,
) string {

	return color +
		text +
		Reset
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
