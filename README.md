# GO_1

# T Test

//Test is _test.go
//type testCase struct {
//	name     string
//	input    int
//	expected int
//}
//
//func TestFibonachiOptim(t *testing.T) {
//	testCase := []testCase{
//		{"case 10", 10, 55},
//		{"case 9", 9, 34},
//		{"case 1", 1, 1},
//	}
//
//	for _, cse := range testCase {
//		cse := cse
//		t.Run(cse.name, func(t *testing.T) {
//			rezult := fibonachiOptim(cse.input, map[int]int{0: 0, 1: 1})
//
//			if rezult != cse.expected {
//				t.Errorf("For %d --%d, got %d", cse.input, cse.expected, rezult)
//
//			}
//		})
//
//	}
//}

//func TestFibonachiOptim(t *testing.T) {
//	rezult := fibonachiOptim(10, map[int]int{0: 0, 1: 1})
//
//	if rezult != 55 {
//		t.Errorf("For 10 --55, got %d", rezult)
//	}
//}

//


//Func is _.go

//func fibonachi(i int) int {
//	if i < 2 {
//		return i
//	}
//
//	return fibonachi(i-1) + fibonachi(i-2)
//}
//
//func fibonachiOptim(n int, hash map[int]int) int {
//	if val, ok := hash[n]; ok {
//		return val
//	}
//	hash[n] = fibonachiOptim(n-1, hash) + fibonachiOptim(n-2, hash)
//
//	return hash[n]
//}

//package main
//
//import (
//	"fmt"
//)
//
//func Examplefibonachi(i int) int {
//	if i < 2 {
//		return i
//	}
//
//	return Examplefibonachi(i-1) + Examplefibonachi(i-2)
//}
//func main() {
//
//	Examplefibonachi(15)
//	{
//		fmt.Println("For 15", "rezult", Examplefibonachi(15))
//		// Output: 610
//	}
//
//}
