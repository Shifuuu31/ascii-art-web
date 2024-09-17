package source

import (
	"fmt"
	"strings"
)

// function responsible for parsong the template file
func getTemplate(data string, templateName string) map[rune][]string {
	startChar := ' '
	var blocks []string
	if templateName == "thinkertoy" {
		data = data[1:]
		blocks = strings.Split(data, "\r\n\r\n")
	} else {
		blocks = strings.Split(data, "\n\n")
	}
	templateMap := make(map[rune][]string)

	for i, block := range blocks {
		var lines []string
		if templateName == "thinkertoy" {
			lines = strings.Split(block, "\r\n")
		} else {
			lines = strings.Split(block, "\n")
		}
		if len(lines) > 0 {
			char := rune(startChar + rune(i))
			templateMap[char] = lines
		} else {
			fmt.Printf("warning: empty or malformed block at index %d", i)
		}

	}

	return templateMap
}

// ascii art generator function
func GenerateAsciiArt(text string, fontMap map[rune][]string) string {
	var result []string
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		var asciiLines [8]string
		for _, char := range line {
			if charArt, found := fontMap[char]; found {
				for i, asciiLine := range charArt {
					asciiLines[i] += asciiLine
				}
			}
		}
		result = append(result, strings.Join(asciiLines[:], "\n"))
	}

	return strings.Join(result, "\n\n")
}
