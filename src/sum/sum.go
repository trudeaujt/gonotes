package sum

//Sum sums together an array of numbers.
//Note that [5]int and [4]int, etc, are technically different types.
//You will get an error if you try and pass in [4]int, just like if you tried passing in a string.
//For handling any size, we can use slices.
//func Sum(numbers [5]int) int {
//	sum := 0
//	// on each iteration, a for range loop gives you two values: the index and value.
//	// we don't need the index here, so instead of i we write _
//	for _ /*i*/, number := range numbers {
//		sum += number
//	}
//	return sum
//}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

//SumAll variadic function to sum all slices passed in.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for i := 0; i < len(numbersToSum); i++ {
		sums = append(sums, Sum(numbersToSum[i]))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			//syntax for slicing is slice[low:hi]
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
