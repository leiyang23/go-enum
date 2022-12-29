package enum

import "testing"

// ColorEnum 枚举示例
type ColorEnum struct {
	RED      Enum
	BLUE     Enum
	BLACK    Enum
	YELLOW   Enum
	BigGreen Enum
}

func TestMakeEnum(t *testing.T) {
	Color := new(ColorEnum)
	err := MakeEnum(Color)
	if err != nil {
		t.Fail()
	}

	if Color.RED.Name != "RED" {
		t.Fail()
	}

	if Color.BigGreen.Name != "big green" {
		t.Fail()
	}
}

func TestValidate(t *testing.T) {
	Color := new(ColorEnum)
	err := MakeEnum(Color)
	if err != nil {
		t.Fail()
	}

	if ErrorEnum != Validate(Color, 20) {
		t.Fail()
	}
	if Validate(Color, 1).Name != "RED" {
		t.Fail()
	}
}

func TestList(t *testing.T) {
	Color := new(ColorEnum)
	err := MakeEnum(Color)
	if err != nil {
		t.Fail()
	}

	if len(List(Color)) != 5 {
		t.Fail()
	}
}
