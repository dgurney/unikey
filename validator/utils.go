package validator

func parseDecimal(value string) (int64, bool) {
	var result int64
	for i := range value {
		if value[i] < '0' || value[i] > '9' {
			return 0, false
		}
		result = result*10 + int64(value[i]-'0')
	}
	return result, true
}

func digitSum(num int64) int64 {
	var s int64
	for num != 0 {
		digit := num % 10
		s += digit
		num /= 10
	}
	return s
}
