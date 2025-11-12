package main

import "fmt"

func main() {
	str1 := "ABABAB"
	str2 := "ABAB"

	fmt.Println(gcdOfStrings(str1, str2))
}

func gcdOfStrings(str1 string, str2 string) string {
	// Проверка str1+str2 != str2+str1 гарантирует, что обе строки состоят из одного и того же “строительного блока”.
	if str1+str2 != str2+str1 {
		return ""
	}

	// Евклид
	a, b := len(str1), len(str2)
	for b > 0 {
		a, b = b, a%b
	}
	// Ответ
	return str1[:a]
}
