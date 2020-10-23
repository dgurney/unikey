package generator

type keyGenerator interface {
	Generate(output chan string)
}
