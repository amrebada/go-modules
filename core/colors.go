package core

import "fmt"

//colors
const (
	green   = "\033[32m"
	white   = "\033[37m"
	yellow  = "\033[33m"
	red     = "\033[31m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	reset   = "\033[0m"
)

var (
	register = Colorize(green, "[Register]")
	generate = Colorize(blue, "[Generate]")
	starting = Colorize(cyan, "[Starting...]")
	migrate  = Colorize(green, "[Migrate]")
	connect  = Colorize(blue, "[Connect]")
)

func Colorize(color string, text string) string {
	return fmt.Sprintf("%s%s%s", color, text, reset)
}
