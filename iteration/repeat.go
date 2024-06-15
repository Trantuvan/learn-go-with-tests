package iteration

// Repeat word 5 times
func Repeat(word string) (result string) {
	for i := 0; i < 5; i++ {
		result += word
	}
	return result
}
