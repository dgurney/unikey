package generator

// Generate generates a key of the specified type
func Generate(k KeyGenerator, ch chan KeyGenerator) {
	k.Generate(ch)
}
