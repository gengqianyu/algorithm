package fibonacci

func Search() int {

}

func Fib(n int) int {
	if n <= 0 {
		return 0
	}

	if n < 3 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}
