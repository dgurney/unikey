package generator

import "testing"

func TestChicago(t *testing.T) {
	c := ChicagoCredentials{Build: "73g", Site: "889884", Password: "fdaa"}
	if err := c.Generate(); err != nil {
		t.Fatalf("should not receive an error, got %s", err)
	}
	if c.String() != "889884/fdaa6c807" {
		t.Fatalf("expected 889884/fdaa6c807, got %s", c)
	}
}

func TestChicagoRejectsUnknownBuild(t *testing.T) {
	c := ChicagoCredentials{Build: "unknown"}
	if err := c.Generate(); err == nil {
		t.Fatal("expected an invalid build error")
	}
}
