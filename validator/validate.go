package validator

// Validate validates a provided key.
func Validate(key KeyValidator, v chan bool) {
	key.Validate(v)
}
