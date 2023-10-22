package artist

import (
	"os"
	"bufio"
	"fmt"
)

var asciiArt []string

func Hasooni() {
textformat := "standard.txt"
//filename := ""
args := os.Args[1:]
for _, v := range args {
	if v == "shadow" || v == "Shadow"{
		textformat = "shadow.txt"
	}
	if v == "thinkertoy" || v == "Thinkertoy"{
		textformat = "thinkertoy.txt"
	}
	if len(v) > 9 && v[:9] == "--output=" {
		textformat = v[9:]
	}
}

	f, err := os.Open(textformat)
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

func GetAsciiArt() []string {
	return asciiArt
}
