package main

import "fmt"

//import "math"

// func main() {
//     add(1, 2, 3)        // sum = 6
//     add(1, 2, 3, 4)     // sum = 10
//     add(5, 6, 7, 2, 3)  // sum = 23
// }

// func add(numbers ...int){
//     var sum = 0
//     for _, number := range numbers{
//         sum += number
//     }
//     fmt.Println("sum = ", sum)
// }

func main() {
	var a int
	fmt.Print("Введите число от - a: ")
	fmt.Scanln(&a)
	ggg(a)
}

func ggg(a int) {
	if a < 0 || a > 999 {
		fmt.Println("Ошибка аргумента, данное значение не находися в интермале [0, 999].")
	} else {
		fmt.Println("Данное знасение находися в интермале [0, 999].")
	}
}

//	return {
//	  units: numb % 10,
//	  dozens: Math.floor(numb / 10) % 10,
//	  hundreds: Math.floor(numb / 100) % 10,
//	  thousands: Math.floor(numb / 1000),
//	};
