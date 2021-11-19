package parser

import (
	"reflect"
	"testing"
)

func TestParseValidatePlateau(t *testing.T) {
	x, y, err := ParseValidatePlateau("5 5")
	if x != 5 {
		t.Errorf("ParseValidatePlateau(\"5 5\") = %d; want 5", x)
	}
	if y != 5 {
		t.Errorf("ParseValidatePlateau(\"5 5\") = %d; want 5", y)
	}
	if err != nil {
		t.Errorf("ParseValidatePlateau(\"5 5 \") return err:  %v", err)
	}
}

func TestParseValidateRoverPosition(t *testing.T) {
	x, y, dir, err := ParseValidateRoverPosition("1 2 N")
	if x != 1 {
		t.Errorf("ParseValidateRoverPosition(\"1 2 N\") = %d; want 1", x)
	}
	if y != 2 {
		t.Errorf("ParseValidateRoverPosition(\"1 2 N\") = %d; want 2", y)
	}
	if dir != "N" {
		t.Errorf("ParseValidateRoverPosition(\"1 2 N\") = %s; want N", dir)
	}
	if err != nil {
		t.Errorf("ParseValidateRoverPosition(\"1 2 N\") return err:  %v", err)
	}
}

func TestParseValidateRoverInstructions(t *testing.T) {
	input_list, err := ParseValidateRoverInstructions("LMLMLMLMM")
	if !reflect.DeepEqual(input_list, []string{"L", "M", "L", "M", "L", "M", "L", "M", "M"}) {
		t.Errorf("ParseValidateRoverPosition(\"LMLMLMLMM\") = %v;", input_list)
	}
	if err != nil {
		t.Errorf("ParseValidateRoverInstructions(\"LMLMLMLMM\") return err:  %v", err)
	}
}
