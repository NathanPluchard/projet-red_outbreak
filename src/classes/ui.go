package classes

import "strings"


const (
	Reset   = "\033[0m"
	Bold    = "\033[1m"
	Dim     = "\033[2m"
	Italic  = "\033[3m"

	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	BrightRed    = "\033[91m"
	BrightGreen  = "\033[92m"
	BrightYellow = "\033[93m"
	BrightBlue   = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan   = "\033[96m"
	BrightWhite  = "\033[97m"
)

// ratio calcule un pourcentage borné entre 0 et 1.
func ratio(current, max int) float64 {
	if max <= 0 {
		return 0
	}
	r := float64(current) / float64(max)
	if r < 0 {
		r = 0
	}
	if r > 1 {
		r = 1
	}
	return r
}


func bar(current, max, width int, color string) string {
	r := ratio(current, max)
	filled := int(r * float64(width))
	if filled > width {
		filled = width
	}
	full := strings.Repeat("█", filled)
	empty := strings.Repeat("░", width-filled)
	return color + full + Reset + Dim + empty + Reset
}


func HPBar(current, max int) string {
	r := ratio(current, max)
	color := Green
	switch {
	case r <= 0.25:
		color = BrightRed
	case r <= 0.5:
		color = Yellow
	}
	return "[" + bar(current, max, 20, color) + "]"
}


func ManaBar(current, max int) string {
	return "[" + bar(current, max, 20, Cyan) + "]"
}

// ExpBar retourne une barre d'expérience magenta.
func ExpBar(current, max int) string {
	return "[" + bar(current, max, 20, Magenta) + "]"
}


func Separator() string {
	return Dim + strings.Repeat("─", 48) + Reset
}


func TitleBox(title string) string {
	inner := " " + title + " "
	width := len([]rune(inner))
	top := "╔" + strings.Repeat("═", width) + "╗"
	mid := "║" + inner + "║"
	bot := "╚" + strings.Repeat("═", width) + "╝"
	return Bold + BrightCyan + top + "\n" + mid + "\n" + bot + Reset
}


func Colorize(color, text string) string {
	return color + text + Reset
}