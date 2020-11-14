package generator

// KeyGenerator is exactly what it says on the tin.
type KeyGenerator interface {
	Generate(k chan KeyGenerator)
	String() string
}
