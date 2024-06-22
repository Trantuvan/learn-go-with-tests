package array

func Sum(intArr ...int) int {
	sum := 0

	for _, v := range intArr {
		sum += v
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	length := len(numbersToSum)

	if length == 0 {
		return nil
	}
	sum := make([]int, length)

	for i, slice := range numbersToSum {
		sum[i] = Sum(slice...)
	}

	return sum
}
