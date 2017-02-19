package colors

import "github.com/sam701/tcolor"

func Match(s string) string {
	return tcolor.Colorize(s, tcolor.New().Underline().ForegroundRGB(4, 1, 1))
}

func Timestamp(s string) string {
	return tcolor.Colorize(s, tcolor.New().Foreground(tcolor.BrightGreen))
}

var PropertyColor = tcolor.New().ForegroundGray(13).Italic()

func Property(s string) string {
	return tcolor.Colorize(s, PropertyColor)
}
