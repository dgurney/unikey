package generator

import (
	"testing"
)

func TestChicago(t *testing.T) {
	c := ChicagoCredentials{Build: "73g", Site: "889884", Password: "fdaa"}
	k, err := Generate(c)
	if err != nil {
		t.Fatalf("should not receive an error, got %s", err)
	}
	if k.String() != "889884/fdaa6c807" {
		t.Fatalf("expected 889884/fdaa6c807, got %s", k)
	}
	t.Log(k.String())
}
