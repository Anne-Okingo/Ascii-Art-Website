package asciiArt

import (
	"bufio"
	"fmt"
	"os"
	"ascii-art-web/banner"
)

// LoadBannerMap loads the banner map from the file provided
func LoadBannerMap(fileName string) (map[int][]string, error) {
	banner.ValidateBanner(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	bannerMap := make(map[int][]string)
	key := 32
	lineCount := 0
	chunk := []string{}

	for scanner.Scan() {
		lines := scanner.Text()

		if lines != "" {
			chunk = append(chunk, lines)
			lineCount++
		}

		if lineCount == 8 {
			bannerMap[key] = chunk
			key++
			chunk = []string{}
			lineCount = 0
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return bannerMap, nil
}
