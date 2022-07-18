package main

import (
	"fmt"
	
)

// func fibonachi(i int) (int) {
// 	if i == 0 {
// 		return 0
// 	}
// 	if i == 1 {
// 		return 1
// 	}
// 	return fibonachi(i-1) + fibonachi(i-2)
// }
// func main() {
// 	var a int
// 	fmt.Print("Введите количество элементов последовательности от 0 до 50: ")
// 	fmt.Scanln(&a)
// 	if a > 0 && a <= 50 {
// 	var i int
// 	for i = 0; i < a; i++ {
// 				fmt.Printf("%d ", fibonachi(i))
// 	}
// 	} else {fmt.Println("Введенное значение некорректно. Попробуйте еще раз.")}
// }
func fib(n int, hash map [int]int) int {
	if val, ok := hash[n]; ok{
		return val
	}
	hash [n] = fib(n-1, hash) + fib(n-2, hash)
	fmt.Println(n, hash[n])
	return hash[n]
	
}

func main () {

	hash := map[int]int{
		0:0,
		1:1,
	}
fmt.Println((fib(92, hash)))


}




