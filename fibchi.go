package fibchi

func fibonachi(i int) (int) {
	if i < 2 {
		return i
	}
	
	return fibonachi(i-1) + fibonachi(i-2)
}

func fibonachiOptim(n int, hash map[int]int) int {
	if val, ok := hash[n]; ok {
		return val } 
	hash[n] = fibonachiOptim(n-1, hash)	+ fibonachiOptim(n-2, hash)

	return hash[n]
}	 