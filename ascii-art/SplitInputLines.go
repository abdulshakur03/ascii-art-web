package ascii_art

import "strings"

func SplitInputLines(text string) []string {
	if text == "" {
		return []string{}
	}

	if text == "\n" {
		return []string{""}
	}

	return strings.Split(text, "\n")
}
