package generator

import (
	"testing"

	"github.com/dgurney/unikey/validator"
)

func Benchmark10digit100(b *testing.B) {
	cd := Mod7CD{}
	kch := make(chan KeyGenerator)
	for n := 0; n < b.N; n++ {
		for i := 0; i < 100; i++ {
			go Generate(cd, kch)
			<-kch
		}
	}
}

func Benchmark11digit100(b *testing.B) {
	cd := Mod7ElevenCD{}
	kch := make(chan KeyGenerator)
	for n := 0; n < b.N; n++ {
		for i := 0; i < 100; i++ {
			go Generate(cd, kch)
			<-kch
		}
	}
}

func BenchmarkOEM100(b *testing.B) {
	cd := Mod7OEM{}
	kch := make(chan KeyGenerator)
	for n := 0; n < b.N; n++ {
		for i := 0; i < 100; i++ {
			go Generate(cd, kch)
			<-kch
		}
	}
}

func TestCD(t *testing.T) {
	cd := Mod7CD{}
	ka := make([]validator.Mod7CD, 0)
	kch := make(chan KeyGenerator)
	vch := make(chan bool)
	for i := 0; i < 500000; i++ {
		go Generate(cd, kch)
		k := <-kch
		ka = append(ka, validator.Mod7CD{k.String()})
	}
	for _, k := range ka {
		t.Log("Validating", k.Key)
		go validator.Validate(k, vch)
		if !<-vch {
			t.Errorf("Generated key %s is invalid!", k.Key)
		}

	}
}

func TestMod7ElevenCD(t *testing.T) {
	cd := Mod7ElevenCD{}
	ka := make([]validator.Mod7ElevenCD, 0)
	kch := make(chan KeyGenerator)
	vch := make(chan bool)
	for i := 0; i < 500000; i++ {
		go Generate(cd, kch)
		k := <-kch
		ka = append(ka, validator.Mod7ElevenCD{k.String()})
	}
	for _, k := range ka {
		t.Log("Validating", k.Key)
		go k.Validate(vch)
		if !<-vch {
			t.Errorf("Generated key %s is invalid!", k.Key)
		}

	}
}

func TestOEM(t *testing.T) {
	cd := Mod7OEM{}
	ka := make([]validator.Mod7OEM, 0)
	kch := make(chan KeyGenerator)
	vch := make(chan bool)
	for i := 0; i < 500000; i++ {
		go Generate(cd, kch)
		k := <-kch
		ka = append(ka, validator.Mod7OEM{k.String()})
	}
	for _, k := range ka {
		t.Log("Validating", k.Key)
		go k.Validate(vch)
		if !<-vch {
			t.Errorf("Generated key %s is invalid!", k.Key)
		}

	}
}
