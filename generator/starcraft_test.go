package generator

import (
	"testing"

	"github.com/dgurney/unikey/validator"
)

func BenchmarkStarCraft(b *testing.B) {
	k := StarCraft{}
	for n := 0; n < b.N; n++ {
		for i := 0; i < 100; i++ {
			k.Generate()
		}
	}
}

func TestStarCraft(t *testing.T) {
	k := StarCraft{}
	for i := 0; i < 500000; i++ {
		k.Generate()
		v := validator.StarCraft{k.String()}
		err := v.Validate()
		if err != nil {
			t.Errorf("invalid key %s generated (%s)", k.String(), err)
			return
		}
	}
}
