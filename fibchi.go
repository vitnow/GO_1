package fibchi

func fibonachi(i int) (int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonachi(i-1) + fibonachi(i-2)
}