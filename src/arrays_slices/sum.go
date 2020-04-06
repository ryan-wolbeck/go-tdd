package arrays

//Sum : Sums all the  numbers in an array
func Sum(nums []int) int{
	sum := 0
	for _, number := range nums {
		sum += number
	}
	return sum
}

//SumAll : Takes in multiple arrays and returns the sums
func SumAll(numbersToSum ...[]int) []int{
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)
	
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}