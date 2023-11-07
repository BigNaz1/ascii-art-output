package main

import (
	//"bufio"
	"fmt"
	"os"
	"strings"
	"artist"
)

func main() {
	// Read command-line arguments
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Insufficient command-line arguments provided.")
		fmt.Println("Usage: go run . [--output=banner.txt] \"Hello There!\" shadow")
		return
	}

	outputFlag := ""           // Stores the output file name
	fontStyle := args[len(args)-1] // Last argument is the font style
	inputText := ""            // Stores the input text to convert to ASCII art

	// Check if the first argument is the --output flag
	if strings.HasPrefix(args[0], "--output=") {
		outputFlag = strings.TrimPrefix(args[0], "--output=")
		inputText = strings.Join(args[1:len(args)-1], " ") // Join the text arguments
	} else {
		inputText = strings.Join(args, " ") // Join all arguments as text
	}

	// Load ASCII art from file
	artist.Hasooni(fontStyle)

	// Generate ASCII art
	asciiArt := artist.generateArt(inputText) // Use the function from the artist package

	// Save the ASCII art to file if the output flag is provided
	if outputFlag != "" {
		err := artist.SaveToFile(outputFlag, asciiArt) // Use the function from the artist package
		if err != nil {
			fmt.Println("Error saving ASCII art:", err)
			return
		}
		fmt.Println("ASCII art saved to", outputFlag)
	} else {
		fmt.Println(asciiArt)
	}
}
