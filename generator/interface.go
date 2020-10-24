package generator

type KeyGenerator interface {
	Generate(k chan KeyGenerator)
	String() string
}
