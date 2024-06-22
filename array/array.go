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

func SumAllTails(numbersToSum ...[]int) []int {
	if len(numbersToSum) == 0 {
		return nil
	}
	sum := []int{}

	for _, slice := range numbersToSum {
		if len(slice) == 0 {
			sum = append(sum, 0)
			continue
		}
		sum = append(sum, Sum(slice[1:]...))
	}
	return sum
}
