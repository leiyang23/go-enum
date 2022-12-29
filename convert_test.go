package enum

import "testing"

func TestCamel2Space(t *testing.T) {

	if Camel2Space("RED") != "RED" {
		t.Fail()
	}

	if Camel2Space("BigGreen") != "big green" {
		t.Fail()
	}
}
