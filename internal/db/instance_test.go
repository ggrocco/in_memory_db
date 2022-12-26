package db

import (
	"testing"
)

func TestInstance(t *testing.T) {
	instance := New()

	instance.Set("test", "123")
	instance.Set("test2", "123")
	instance.Set("test3", "123")

	test := instance.Get("test")
	if test != "123" {
		t.Error("Fail on get test value")
	}

	invalid := instance.Get("invalid")
	if invalid != "Nil" {
		t.Error("Fail on get Nil for invalid")
	}

	numEqualTo := instance.NumEqualTo("123")
	if numEqualTo != 3 {
		t.Error("Fail on NumEqualTo")
	}

	invalidNumEqualTo := instance.NumEqualTo("2")
	if invalidNumEqualTo != 0 {
		t.Error("Fail on NumEqualTo")
	}

	instance.Unset("test")

	test = instance.Get("test")
	if test != "Nil" {
		t.Error("Fail on get Nil for test after unset")
	}

	numEqualTo = instance.NumEqualTo("123")
	if numEqualTo != 2 {
		t.Error("Fail on NumEqualTo after unset")
	}
}
