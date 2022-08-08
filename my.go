package my


//Exercise_1
//Benchmark -Сортировка вставкой

func Insertsort(ars []int) {
	for i := 1; i < len(ars); i++ {
		x := ars[i]
		j := i
		for ; j >= 1 && ars[j-1] > x; j-- {
			ars[j] = ars[j-1]
		}
		ars[j] = x
	}
}

// Benchmark - Sortfast сортировка

func Quicksort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	Quicksort(ar[:split])
	Quicksort(ar[split:])
}

func partition(ar []int) int {
	pivot := ar[len(ar)/2]

	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < pivot; left++ {
		}

		for ; ar[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		swap(ar, left, right)
	}
}

func swap(ar []int, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
}
//Exercise_2
// Табличный тест Метод Фибоначи

func fibonachiOptim(n int, hash map[int]int) int {
	if val, ok := hash[n]; ok {
		return val
	}
	hash[n] = fibonachiOptim(n-1, hash) + fibonachiOptim(n-2, hash)

	return hash[n]
}
//Exercise_3
// Example test

func Examplefibonachi(i int) int {
	if i < 2 {
		return i
	}

	return Examplefibonachi(i-1) + Examplefibonachi(i-2)
}