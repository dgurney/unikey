package validator

func digitsum(num int64) int64 {
	var s int64
	for num != 0 {
		digit := num % 10
		s += digit
		num /= 10
	}
	return s
}

func checkdigitCheck(k int64) bool {
	// Check digit cannot be 0 or >= 8.
	if k%10 == 0 || k%10 >= 8 {
		return false
	}
	return true
}
