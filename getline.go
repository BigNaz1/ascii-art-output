package artist

func Getline(num int) string {
	if num < len(asciiArt) {
		return asciiArt[num]
	}
	return ""
}
