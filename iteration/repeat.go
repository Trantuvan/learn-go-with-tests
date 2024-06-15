package iteration

// Repeat word n times
func Repeat(word string, n int) (result string) {
	for i := 0; i < n; i++ {
		result += word
	}
	return result
}
