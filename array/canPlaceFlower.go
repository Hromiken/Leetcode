package main

import "fmt"

func main() {
	s := make([]int, 5)
	s[0] = 0
	s[1] = 0
	s[2] = 1
	s[3] = 0
	s[4] = 1
	fmt.Println(canPlaceFlowers(s, 1))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	cnt := 0
	for i, _ := range flowerbed {
		// Проверка есть ли цветок в ячейке
		if flowerbed[i] != 0 {
			continue
		}

		// Если это не первая позиция и слева есть цветок, переходим к следующей.
		if i != 0 && flowerbed[i-1] != 0 {
			continue
		}

		// Если это не последняя позиция и справа есть цветок, переходим к следующей.
		if i != len(flowerbed)-1 && flowerbed[i+1] != 0 {
			continue
		}

		// Если все проверки пройдены, сажаем цветок в текущую позицию и увеличиваем счетчик cnt.
		flowerbed[i] = 1
		cnt++
	}
	if cnt >= n {
		return true
	}
	return false
}
