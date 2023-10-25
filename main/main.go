package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"artist"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Insufficient command-line arguments provided.")
		fmt.Println("Usage: go run . [--output=banner.txt] \"Hello There!\" shadow")
		return
	}

	outputFlag := ""
	fontStyle := args[len(args)-1]
	inputText := ""

	// Check if the first argument is the --output flag
	if strings.HasPrefix(args[0], "--output=") {
		outputFlag = strings.TrimPrefix(args[0], "--output=")
		inputText = strings.Join(args[1:len(args)-1], " ")
	} else {
		inputText = strings.Join(args, " ")
	}

	// Load ASCII art from file
	artist.Hasooni(fontStyle)

	// Print the ASCII art
	asciiArt := generateASCIIArt(inputText)

	// Save the ASCII art to file if the output flag is provided
	if outputFlag != "" {
		err := saveToFile(outputFlag, asciiArt)
		if err != nil {
			fmt.Println("Error saving ASCII art:", err)
			return
		}
		fmt.Println("ASCII art saved to", outputFlag)
	} else {
		fmt.Println(asciiArt)
	}
}

func generateASCIIArt(text string) string {
	lines := strings.Split(text, "\\n")
	var asciiArt strings.Builder

	for _, line := range lines {
		for i := 0; i < 8; i++ {
			for _, letter := range line {
				asciiArt.WriteString(artist.Getline(1 + int(letter-' ')*9 + i))
			}
			asciiArt.WriteString("\n")
		}
		asciiArt.WriteString("\n")
	}

	return asciiArt.String()
}

func saveToFile(fileName, content string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

