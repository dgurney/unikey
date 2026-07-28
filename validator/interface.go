package validator

// KeyValidator validates a license key.
type KeyValidator interface {
	Validate() error
}
