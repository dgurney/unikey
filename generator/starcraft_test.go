package generator

import (
	"testing"

	"github.com/dgurney/unikey/validator"
)

func BenchmarkStarCraft(b *testing.B) {
	k := StarCraft{}
	for range b.N {
		for range 100 {
			k.Generate()
		}
	}
}

func TestStarCraft(t *testing.T) {
	k := StarCraft{}
	for range 10_000 {
		k.Generate()
		v := validator.StarCraft{Key: k.String()}
		if err := v.Validate(); err != nil {
			t.Fatalf("generated key %s is invalid: %v", k, err)
		}
	}
}

func TestStarCraftZeroValue(t *testing.T) {
	var key StarCraft
	if key.String() != "" {
		t.Fatalf("zero-value key should be empty, got %q", key)
	}
}
