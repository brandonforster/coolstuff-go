package stringstuff

func RemoveChars(str string, toRemove string) string {
	ignore := make(map[rune]bool)

	for _, char := range toRemove {
		ignore[char] = true
	}

	output := ""
	for _, char := range str {
		if !ignore[char] {
			output = output + string(char)
		}
	}

	return output
}
