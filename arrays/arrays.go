package main

// func Sum(numbers [5]int) int {
// 	sum := 0

// 	// commented out because of refactored code
// 	// for i := 0; i < 5; i++ {
// 	// 	sum += numbers[i]
// 	// }

// 	// range lets you iterate over an array. Every time it is called it returns two values, the index and the value.
// 	// We are choosing to ignore the index value by using _ blank identifier.
// 	for _, number := range numbers {
// 		sum += number
// 	}

// 	return sum
// }

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// func SumAll(arr1, arr2 []int) []int {
// 	return []int{3, 9}
// }

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }

	// refactor
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
