package generator

func digitsum(num int) int {
	s := 0
	for num != 0 {
		digit := num % 10
		s += digit
		num /= 10
	}
	return s
}
