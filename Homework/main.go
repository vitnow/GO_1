package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a, b, am int     					     // Запрос значений у пользователя
	fmt.Print("Введите минимальное значение min: ")
	fmt.Scanln(&a)
	fmt.Print("Введите минимальное значение max: ")
	fmt.Scanln(&b)
	fmt.Print("Введите колличество элементов: ")
	fmt.Scanln(&am)
	ars := make([]int, am)
	min := a
	max := b

	for a := 1; a < am; a++ {                    //Генерация  и запись в массив сгенерированные значения 
		ars[a] = (rand.Intn(max-min) + min)
		
	}
	
	fmt.Println("Сгенерированный ряд значенией", ars)
	
	for i := 1; i < len(ars); i++ { 			//Сортировка значений методом --Вставка
		x := ars[i]
		j := i
		for ; j >= 1 && ars[j-1] > x; j-- {
			ars[j] = ars[j-1]
		}
		ars[j] = x
	}
	fmt.Println("Отсартированный ряд значенией",ars)
}
