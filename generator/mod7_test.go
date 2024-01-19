package generator

import (
	"fmt"
	"testing"

	"github.com/dgurney/unikey/validator"
)

func Benchmark10digit100(b *testing.B) {
	cd := Mod7CD{}
	for n := 0; n < b.N; n++ {
		for i := 0; i < 100; i++ {
			cd.Generate()
		}
	}
}

func Benchmark11digit100(b *testing.B) {
	cd := Mod7ElevenCD{}
	for n := 0; n < b.N; n++ {
		for i := 0; i < 100; i++ {
			cd.Generate()
		}
	}
}

func BenchmarkOEM100(b *testing.B) {
	cd := Mod7OEM{}
	for n := 0; n < b.N; n++ {
		for i := 0; i < 100; i++ {
			cd.Generate()
		}
	}
}

func TestCD(t *testing.T) {
	t.Parallel()
	cd := Mod7CD{}
	ka := make([]validator.Mod7CD, 0)
	for i := 0; i < 500000; i++ {
		cd.Generate()
		ka = append(ka, validator.Mod7CD{First: fmt.Sprintf("%03d", cd.First), Second: fmt.Sprintf("%07d", cd.Second), Is95: false})
	}
	for _, k := range ka {
		err := k.Validate()
		if err != nil {
			t.Errorf("Generated key %s-%s is invalid!", k.First, k.Second)
		}

	}
}

func TestMod7ElevenCD(t *testing.T) {
	t.Parallel()
	cd := Mod7ElevenCD{}
	ka := make([]validator.Mod7ElevenCD, 0)
	for i := 0; i < 500000; i++ {
		cd.Generate()
		ka = append(ka, validator.Mod7ElevenCD{First: fmt.Sprintf("%04d", cd.First), Second: fmt.Sprintf("%07d", cd.Second)})
	}
	for _, k := range ka {
		err := k.Validate()
		if err != nil {
			t.Errorf("Generated key %s-%s is invalid!", k.First, k.Second)
		}

	}
}

func TestOEM(t *testing.T) {
	t.Parallel()
	oem := Mod7OEM{}
	ka := make([]validator.Mod7OEM, 0)
	for i := 0; i < 500000; i++ {
		oem.Generate()
		ka = append(ka, validator.Mod7OEM{First: oem.First, Second: oem.Second, Third: fmt.Sprintf("%07d", oem.Third), Fourth: fmt.Sprintf("%05d", oem.Fourth), Is95: false})
	}
	for _, k := range ka {
		err := k.Validate()
		if err != nil {
			t.Errorf("Generated key %s-%s-%s-%s is invalid!", k.First, k.Second, k.Third, k.Fourth)
		}

	}
}
