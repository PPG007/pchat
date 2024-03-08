package utils

import "github.com/fatih/color"

func GetColorString(src string, attribute color.Attribute) string {
	return color.New(attribute).Sprintf(src)
}

func ColorStringFn(attribute color.Attribute) func(string) string {
	fn := color.New(attribute).SprintfFunc()
	return func(s string) string {
		return fn(s)
	}
}
