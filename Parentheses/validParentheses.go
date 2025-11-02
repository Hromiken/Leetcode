package main

import "fmt"

func main() {
	input := "()[]{}"
	fmt.Println(isValid(input))
}

func isValid(s string) bool {
	var stack []rune
	pairs := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	for _, r := range s {
		_, ok := pairs[r]
		if ok {
			stack = append(stack, r)
			continue
		}
		// прошли мимо if-a сверху
		// мы имеем дело с закрытой скобкой

		// если прошли по условию у нас пустой стек и закрытая скобка, дальше идти нет смысла
		if len(stack) == 0 {
			return false
		}
		// pop operation
		lastChar := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// проверяем что последний элемент в стеке и текущий образуют пару
		// если они не образуют пару то вернем false
		if a := pairs[lastChar]; a != r {
			return false
		}
	}
	return len(stack) == 0
}
