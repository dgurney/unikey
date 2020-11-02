package generator

import (
	"fmt"
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
		c, _ := k.(Mod7CD)
		ka = append(ka, validator.Mod7CD{fmt.Sprintf("%03d", c.First), fmt.Sprintf("%07d", c.Second)})
	}
	for _, k := range ka {
		t.Logf("Validating %s-%s", k.First, k.Second)
		go validator.Validate(k, vch)
		if !<-vch {
			t.Errorf("Generated key %s-%s is invalid!", k.First, k.Second)
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
		e := k.(Mod7ElevenCD)
		ka = append(ka, validator.Mod7ElevenCD{e.First, fmt.Sprintf("%07d", e.Second)})
	}
	for _, k := range ka {
		t.Logf("Validating %s-%s", k.First, k.Second)
		go k.Validate(vch)
		if !<-vch {
			t.Errorf("Generated key %s-%s is invalid!", k.First, k.Second)
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
		o := k.(Mod7OEM)
		ka = append(ka, validator.Mod7OEM{o.First, o.Second, fmt.Sprintf("%07d", o.Third), fmt.Sprintf("%05d", o.Fourth)})
	}
	for _, k := range ka {
		t.Logf("Validating %s-%s-%s-%s", k.First, k.Second, k.Third, k.Fourth)
		go k.Validate(vch)
		if !<-vch {
			t.Errorf("Generated key %s-%s-%s-%s is invalid!", k.First, k.Second, k.Third, k.Fourth)
		}

	}
}
