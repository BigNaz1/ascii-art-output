package artist

import (
	"bufio"
	"fmt"
	"os"
)

var asciiArt []string

func Hasooni(font string) {
	textFormat := "standard.txt"

	switch font {
	case "shadow":
		textFormat = "shadow.txt"
	case "thinkertoy":
		textFormat = "thinkertoy.txt"
	}

	f, err := os.Open(textFormat)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		asciiArt = append(asciiArt, scanner.Text())
	}
}
