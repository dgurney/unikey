package generator

import (
	"fmt"
	"testing"

	"github.com/dgurney/unikey/validator"
)

func Benchmark10digit100(b *testing.B) {
	cd := Mod7CD{}
	for range b.N {
		for range 100 {
			cd.Generate()
		}
	}
}

func Benchmark11digit100(b *testing.B) {
	cd := Mod7ElevenCD{}
	for range b.N {
		for range 100 {
			cd.Generate()
		}
	}
}

func BenchmarkOEM100(b *testing.B) {
	cd := Mod7OEM{}
	for range b.N {
		for range 100 {
			cd.Generate()
		}
	}
}

func TestCD(t *testing.T) {
	t.Parallel()
	cd := Mod7CD{}
	for range 10_000 {
		cd.Generate()
		key := validator.Mod7CD{
			First:  fmt.Sprintf("%03d", cd.First),
			Second: fmt.Sprintf("%07d", cd.Second),
		}
		if err := key.Validate(); err != nil {
			t.Fatalf("generated key %s is invalid: %v", cd, err)
		}
	}
}

func TestMod7ElevenCD(t *testing.T) {
	t.Parallel()
	cd := Mod7ElevenCD{}
	for range 10_000 {
		cd.Generate()
		key := validator.Mod7ElevenCD{
			First:                fmt.Sprintf("%04d", cd.First),
			Second:               fmt.Sprintf("%07d", cd.Second),
			EnableCheckDigitRule: true,
		}
		if err := key.Validate(); err != nil {
			t.Fatalf("generated key %s is invalid: %v", cd, err)
		}
	}
}

func TestOEM(t *testing.T) {
	t.Parallel()
	oem := Mod7OEM{}
	for range 10_000 {
		oem.Generate()
		key := validator.Mod7OEM{
			First:  oem.First,
			Second: oem.Second,
			Third:  fmt.Sprintf("%07d", oem.Third),
			Fourth: fmt.Sprintf("%05d", oem.Fourth),
		}
		if err := key.Validate(); err != nil {
			t.Fatalf("generated key %s is invalid: %v", oem, err)
		}
	}
}
