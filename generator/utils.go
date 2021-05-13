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

func isLeap(year int) bool {
	switch {
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	case year%4 == 0:
		return true
	default:
		return false
	}
}
