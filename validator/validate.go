package validator

// Validate validates a provided key.
func Validate(key KeyValidator) error {
	return key.Validate()
}
