package generator

import (
	"testing"
)

func TestChicago(t *testing.T) {
	c := ChicagoCredentials{Build: "73g", Site: "889884", Password: "fdaa"}
	err := c.Generate()
	if err != nil {
		t.Fatalf("should not receive an error, got %s", err)
	}
	if c.String() != "889884/fdaa6c807" {
		t.Fatalf("expected 889884/fdaa6c807, got %s", c)
	}
	t.Log(c.String())
}
