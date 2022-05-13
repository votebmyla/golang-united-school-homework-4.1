package reverse_string

func ReverseString(input string) (output string) {
	runeSlice := []rune(input)
	for i := len(runeSlice); i > 0; i-- {
		runeSlice = append(runeSlice, runeSlice[i-1])
	}
	return string(runeSlice[len(runeSlice)/2:])
}
