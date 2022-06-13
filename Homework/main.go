package main

import (
	"fmt"
//	"math/rand"
)

func main() {
// for i :=1; i < 40; i++{
// 	min := -100
//     max := 100
// 	fmt.Print(rand.Intn(max - min) + min)
// }
	
	ars := []int{-19, -13, -53, -41, -19, 18, -75, 40, -44, 0, -6, 11, 62, -11, 28, -26, -89, -55, -63, 6, -5, -34, 28, -42, -53, 47,
	13, -12, 90, -85, 41, -92, 87, -69, -71, 56, 37, -69, -15}
	SortInsert(ars)
	fmt.Println(ars)
}

func SortInsert(ars []int) {
	for i := 1; i < len(ars); i++ {
		x := ars[i]
		j := i
		for ; j >= 1 && ars[j-1] > x; j-- {
			ars[j] = ars[j-1]
		}
		ars[j] = x
	}
}