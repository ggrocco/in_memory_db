package db

import "testing"

func TestMemory(t *testing.T) {
	memory := newMemory()

	memory.Set("test", "123")
	memory.Set("test2", "123")
	memory.Set("test3", "123")

	test := memory.Get("test")
	if test != "123" {
		t.Error("Fail on get test value")
	}

	invalid := memory.Get("invalid")
	if invalid != "Nil" {
		t.Error("Fail on get Nil for invalid")
	}

	numEqualTo := memory.NumEqualTo("123")
	if numEqualTo != 3 {
		t.Error("Fail on NumEqualTo")
	}

	invalidNumEqualTo := memory.NumEqualTo("2")
	if invalidNumEqualTo != 0 {
		t.Error("Fail on NumEqualTo")
	}

	memory.Unset("test")

	test = memory.Get("test")
	if test != "Nil" {
		t.Error("Fail on get Nil for test after unset")
	}

	numEqualTo = memory.NumEqualTo("123")
	if numEqualTo != 2 {
		t.Error("Fail on NumEqualTo after unset")
	}
}
