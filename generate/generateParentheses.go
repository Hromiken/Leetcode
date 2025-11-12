package main

import "fmt"

func main() {
	ex2 := 4

	fmt.Println(generateParenthesis(ex2))

}

func generateParenthesis(n int) []string {
	// Инициализация ф-ии
	var solve func(open, close int, temp string)
	// capacity + ответ
	capAnswer := n
	if capAnswer != 1 && capAnswer != 2 {
		capAnswer = capAnswer + capAnswer - 1
	}
	ans := make([]string, 0, capAnswer)
	// Реализация ф-ии
	solve = func(open, close int, temp string) {
		//open - сколько открывающих скобок ( ещё можно использовать
		//close - сколько закрывающих скобок ) ещё можно использовать
		//temp - текущая формируемая строка

		// Базовый случай
		if len(temp) == 2*n {
			ans = append(ans, temp)
			return
		}

		//Можно добавить ( если у нас ещё есть неиспользованные открывающие скобки
		//Уменьшаем счётчик открывающих скобок на 1
		if open > 0 {
			solve(open-1, close, temp+"(")
		}

		// Ключевое условие! Можно добавить ) только если закрывающих скобок осталось БОЛЬШЕ, чем открывающих
		//Это гарантирует, что у нас всегда будет достаточно открывающих скобок для закрытия
		//Уменьшаем счётчик закрывающих скобок на 1
		if close > open {
			solve(open, close-1, temp+")")
		}

	}
	// Исполнение ф-ии
	solve(n, n, "")
	return ans
}
