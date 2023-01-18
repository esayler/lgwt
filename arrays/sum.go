package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	lengthOfNumbers := len(numbers)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbers {
		sums[i] = Sum(numbers)
	}

	return sums
}
