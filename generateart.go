package artist

import (
	"strings"
)

// generateASCIIArt generates the ASCII art for the given text
func generateASCIIArt(text string) string {
	lines := strings.Split(text, "\\n") // Split text into lines
	var asciiArt strings.Builder

	for _, line := range lines {
		for i := 0; i < 8; i++ { // Iterate over each line and append ASCII characters
			for _, letter := range line {
				asciiArt.WriteString(Getline(1 + int(letter-' ')*9 + i))
			}
			asciiArt.WriteString("\n")
		}
		asciiArt.WriteString("\n")
	}

	return asciiArt.String()
}
