package my

import (
	"fmt"
	"testing"
)


// Exercise_1
//func BenchmarkQuicksort(b *testing.B) {
	ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
	for i := 0; i < b.N; i++ {
		Quicksort(ar)

	}
}
func BenchmarkInsertsort(b *testing.B) {
	ars := []int{3, 4, 1, 2, 5, 7, -1, 0}
	for i := 0; i < b.N; i++ {
		Insertsort(ars)
	}
}

// Exercise_2

type testCase struct {
	name     string
	input    int
	expected int
}


func TestFibonachiOptim(t *testing.T) {
	testCase := []testCase{
		{"case 10", 10, 55},
		{"case 9", 9, 34},
		{"case 1", 1, 1},
	}

	for _, cse := range testCase {
		cse := cse
		t.Run(cse.name, func(t *testing.T) {
			rezult := fibonachiOptim(cse.input, map[int]int{0: 0, 1: 1})

			if rezult != cse.expected {
				t.Errorf("For %d --%d, got %d", cse.input, cse.expected, rezult)

			}
		})

	}
}

// Exercise_3
// Example test run

func Examplefibonachi(15) {
	fmt.Println("For 15", "rezult", Examplefibonachi(15))
	// Output: 610
}

