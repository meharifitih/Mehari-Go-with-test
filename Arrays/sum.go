package main

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func sumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, sum(numbers))
	}
	return sums
}

// func sumAll(numbersToSum ...[]int) []int {
// 	lengthOfNumbers := len(numbersToSum)
// 	sums := make([]int, lengthOfNumbers)

// 	for i, numbers := range numbersToSum {
// 		sums[i] = sum(numbers)
// 	}
// 	return sums
// }

// func Sum(numbers []int) int {
// 	sum := 0
// 	for _, number := range numbers {
// 		sum += number
// 	}
// 	return sum
// }

func main() {
	// fmt.Println(sum([5]int{1, 2, 3, 4, 5}))
}
