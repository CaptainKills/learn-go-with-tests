package sum

func Sum(numbers []int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int

	for _, value := range numbers {
		sums = append(sums, Sum(value))
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int

	for _, value := range numbers {
		if len(value) == 0 {
			sums = append(sums, 0)
			continue
		}

		tail := value[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
