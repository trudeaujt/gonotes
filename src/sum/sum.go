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
	add := func(x, y int) int {
		return x + y
	}
	result := Reduce(numbers, add, 0)
	return result
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
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numbersToSum, sumTail, []int{})
}

func Reduce[A any](collection []A, accumulator func(A, A) A, initialValue A) A {
	var result = initialValue
	for _, a := range collection {
		result = accumulator(result, a)
	}
	return result
}
