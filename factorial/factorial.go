package factorial

/*
func Factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}
*/

func Factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		s := 1
		for i := n; i > 0; i-- {
			s = s * i
		}
		return s
	}
}
