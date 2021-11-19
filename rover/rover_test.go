package rover

import "testing"

func init_test_rover() Rover {
	plateau := Plateau{X_bound: 5, Y_bound: 5}
	position := Position{X: 1, Y: 2}
	mars_rover := Rover{Position: &position,
		Plateau:   plateau,
		Direction: NORTH}

	return mars_rover
}

func TestReadInstructions(t *testing.T) {
	r := init_test_rover()
	err := r.ReadInstructions([]string{"L", "M", "L", "M", "L", "M", "L", "M", "M"})
	if r.Position.X != 1 {
		t.Errorf("ReadInstructions final position X: %d; want 1", r.Position.X)
	}
	if r.Position.Y != 3 {
		t.Errorf("ReadInstructions final position Y: %d; want 3", r.Position.Y)
	}
	if r.Direction != NORTH {
		t.Errorf("ReadInstructions rover final direction: %s; want N", r.Direction)
	}
	if err != nil {
		t.Errorf("ReadInstructions err:  %v", err)
	}

	r = init_test_rover()
	err = r.ReadInstructions([]string{"L", "M", "L", "M", "L", "M", "L", "M", "M", "R"})
	if r.Position.X != 1 {
		t.Errorf("ReadInstructions final position X: %d; want 1", r.Position.X)
	}
	if r.Position.Y != 3 {
		t.Errorf("ReadInstructions final position Y: %d; want 3", r.Position.Y)
	}
	if r.Direction != EAST {
		t.Errorf("ReadInstructions rover final direction: %s; want E", r.Direction)
	}
	if err != nil {
		t.Errorf("ReadInstructions err:  %v", err)
	}
}
