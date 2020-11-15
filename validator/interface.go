package validator

// KeyValidator is exactly what it says on the tin.
type KeyValidator interface {
	Validate() error
}
