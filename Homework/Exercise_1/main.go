package main

import (
	"fmt"
	//"math"
	//	"os"
)

/*  Доработать калькулятор из методички: больше операций и валидации данных
 (проверка деления на 0, возведение в степень, факториал)
+ форматирование строки при выводе дробного числа ( 2 знака после точки)
*/
func factorial(n uint) uint {

	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var a, b, rest float32
	//	var restS float32
	var op string
	fmt.Print("Введите число а: ")
	fmt.Scanln(&a)
	fmt.Print("Введите число b: ")
	fmt.Scanln(&b)
	fmt.Print("Введите арифметическую операцию (+, -, *, /, x2, x3, x^n, n!): ")
	fmt.Scanln(&op)
	if op == "+" || op == "-" || op == "*" {
		if op == "+" {
			rest = a + b
		} else if op == "-" {
			rest = a - b
		} else {
			rest = a * b
		}
		fmt.Printf("Результат выполнения операции: %.2f\n", rest)
	} else if op == "/" && b != 0 {
		rest = a / b
		fmt.Printf("Результат выполнения операции: %.2f\n", rest)
	} else if op == "n!" {
				
		//var n uint64 
		fmt.Println("Результат выполнения операции: ", factorial(6)) // 24
		}  else { fmt.Println("Операция выбрана неверно")
}

//	switch op {
//	case "+":
//		rest = a + b
//	case "-":
//		rest = a - b
//	case "*":
//		rest = a * b
//	case "/":
//		rest = a / b
//	case "x2":
//		rest = a * a
//		restS = b * b
//	case "x3":
//		rest = a * a * a
//		restS = b * b * b
//	case "x^n":
//		var x, k, y float32
//		y = a
//		k = b
//		for k = 1; k < (k - 1); k++ {
//			x = x * y
//		}
//		rest = x
//	//case "n!":
//	//	rest = ((a − 2) * (a − 1)) * a
//	default:
//		fmt.Println("Операция выбрана неверно")
//		os.Exit(1)
//	}
//	fmt.Printf("Результат выполнения операции: %.2f\n", rest)
//	fmt.Printf("Результат выполнения операции: %.2f\n", restS)
//}
}
