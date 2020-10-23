package validator

type keyValidator interface {
	Validate(key string, valid chan bool)
}
