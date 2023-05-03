package generator

// KeyGenerator is exactly what it says on the tin.
type KeyGenerator interface {
	Generate() error
	String() string
}
