package main

import "fmt"

func main() {
	array := make([]int, 5)
	array[0] = 2
	array[1] = 3
	array[2] = 5
	array[3] = 1
	array[4] = 3
	extra := 3
	fmt.Println(kidsWithCandies(array, extra))

}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	answer := make([]bool, len(candies))
	// находим максимальное кол-во конфет (без extra у всех детей)
	maxCand := 0
	for i := 0; i < len(candies); i++ {
		if candies[i] > maxCand {
			maxCand = candies[i]
		}
	}
	// если с extra даже не хватает перебить maxCand => false
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= maxCand {
			answer[i] = true
		}
	}
	return answer
}
