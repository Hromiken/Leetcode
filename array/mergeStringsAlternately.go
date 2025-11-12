package main

import "fmt"

func main() {
	word1 := "abc"
	word2 := "pqr"
	fmt.Println(mergeAlternately(word1, word2))

}

func mergeAlternately(word1 string, word2 string) string {
	ans1 := []rune(word1)
	ans2 := []rune(word2)
	answer := make([]rune, 0, len(ans1)+len(ans2))

	// находим максимальную длину
	maxLen := len(ans1)
	if len(ans2) > maxLen {
		maxLen = len(ans2)
	}

	for i := 0; i < maxLen; i++ {
		// Добавляем из первого слова, если есть
		if i < len(ans1) {
			answer = append(answer, ans1[i])
		}
		// Добавляем из второго слова, если есть
		if i < len(ans2) {
			answer = append(answer, ans2[i])
		}
	}
	return string(answer)
}
