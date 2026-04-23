package ascii_art

import (
	"os"
	"strings"
)

func ReadBannerFile(fileName string) ([]string, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")

	banner := strings.Split(content, "\n")

	return banner, nil

}
