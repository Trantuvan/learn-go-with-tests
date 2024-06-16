package array

func Sum(intArr ...int) int {
	sum := 0

	for _, v := range intArr {
		sum += v
	}
	return sum
}
