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
		ka = append(ka, validator.Mod7CD{fmt.Sprintf("%03d", cd.First), fmt.Sprintf("%07d", cd.Second), false})
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
		ka = append(ka, validator.Mod7ElevenCD{fmt.Sprintf("%04d", cd.First), fmt.Sprintf("%07d", cd.Second)})
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
		ka = append(ka, validator.Mod7OEM{oem.First, oem.Second, fmt.Sprintf("%07d", oem.Third), fmt.Sprintf("%05d", oem.Fourth), false})
	}
	for _, k := range ka {
		err := k.Validate()
		if err != nil {
			t.Errorf("Generated key %s-%s-%s-%s is invalid!", k.First, k.Second, k.Third, k.Fourth)
		}

	}
}
