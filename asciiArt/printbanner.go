package asciiArt

import (
	"fmt"
	"strings"
)

// Print the banner from a line of text
func PrintLineBanner(line string, bannerMap map[int][]string) (string, error) {
	var lineoutput []string
	args := strings.ReplaceAll(line, "\r", "\n")
	lines := strings.Split(args, "\n")
	for _, word := range lines {
		if word == "" {
			lineoutput = append(lineoutput, "\n")
			continue
		}

		output := make([]string, 8)

		for _, char := range word {
			banner, exists := bannerMap[int(char)]
			if !exists {
				return "", fmt.Errorf("character %c not found in banner map", char)
			}

			for i := 0; i < 8; i++ {
				output[i] += banner[i]
			}
		}
		lineoutput = append(lineoutput, strings.Join(output, "\n"))
	}

	return strings.Join(lineoutput, "\n\n"), nil
}
