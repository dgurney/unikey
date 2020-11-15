package generator

// Generate generates a key of the specified type
func Generate(k KeyGenerator) (KeyGenerator, error) {
	key, err := k.Generate()
	return key, err
}
