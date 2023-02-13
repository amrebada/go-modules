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

func GetStatusString(status HttpMethods) string {
	switch status {
	case HTTP_GET_METHOD:
		return Colorize(blue, fmt.Sprintf("[%s]", string(HTTP_GET_METHOD)))
	case HTTP_POST_METHOD:
		return Colorize(cyan, fmt.Sprintf("[%s]", string(HTTP_POST_METHOD)))
	case HTTP_DELETE_METHOD:
		return Colorize(red, fmt.Sprintf("[%s]", string(HTTP_DELETE_METHOD)))
	case HTTP_PUT_METHOD:
		return Colorize(yellow, fmt.Sprintf("[%s]", string(HTTP_PUT_METHOD)))
	case HTTP_PATCH_METHOD:
		return Colorize(green, fmt.Sprintf("[%s]", string(HTTP_PATCH_METHOD)))
	default:
		return string(status)
	}
}
