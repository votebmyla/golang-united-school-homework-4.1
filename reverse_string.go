package reverse_string

func ReverseString(input string) (output string) {
	byteSlice := []byte(input)
	for i := len(byteSlice); i > 0; i-- {
		byteSlice = append(byteSlice, byteSlice[i-1])
	}
	return string(byteSlice[len(byteSlice)/2:])
}
