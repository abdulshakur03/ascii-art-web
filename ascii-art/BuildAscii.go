package ascii_art

import "strings"

func BuildAscii(lines []string, banner []string) string {
	var result strings.Builder

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}

		for row := 1; row < 9; row++ {
			for _, char := range line {
				if char < 32 || char > 126 {
					continue
				}
				index := (int(char) - 32) * 9

				result.WriteString(banner[index+row])

			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
