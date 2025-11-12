package main

import "fmt"

func main() {
	nums := make([]int, 3)
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	fmt.Println(getConcatenation(nums))
}

func getConcatenation(nums []int) []int {
	answer := make([]int, 0, len(nums)*2)
	answer = append(answer, nums...)
	answer = append(answer, nums...)
	return answer
}
