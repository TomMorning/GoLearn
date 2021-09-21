package array

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers) // make创建一个指定大小的新的slice
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}
