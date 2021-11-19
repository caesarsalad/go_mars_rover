package rover

import "errors"

const (
	NORTH = "N"
	EAST  = "E"
	SOUTH = "S"
	WEST  = "W"
	RIGHT = "R"
	LEFT  = "L"
	MOVE  = "M"
)

var (
	ErrReachBound = errors.New("rover reached the plateau bounds")
)

type Position struct {
	X int
	Y int
}

type Plateau struct {
	X_bound int
	Y_bound int
}

type Rover struct {
	Position  *Position
	Direction string
	Plateau   Plateau
}

func (r *Rover) ReadInstructions(instructions []string) error {
	var err error
	for _, instruction := range instructions {
		switch instruction {
		case MOVE:
			err = r.Move()
			if err != nil {
				return err
			}
		case RIGHT:
			r.TurnRight()
		case LEFT:
			r.TurnLeft()
		}
	}
	return nil
}

func (r *Rover) Move() error {
	switch r.Direction {
	case NORTH:
		if r.Plateau.Y_bound < r.Position.Y+1 {
			return ErrReachBound
		}
		r.Position.Y += 1
	case EAST:
		if r.Plateau.X_bound < r.Position.X+1 {
			return ErrReachBound
		}
		r.Position.X += 1
	case SOUTH:
		if r.Position.Y-1 < 0 {
			return ErrReachBound
		}
		r.Position.Y -= 1
	case WEST:
		if r.Position.X-1 < 0 {
			return ErrReachBound
		}
		r.Position.X -= 1
	}
	return nil
}

func (r *Rover) TurnRight() {
	switch r.Direction {
	case NORTH:
		r.Direction = EAST
	case EAST:
		r.Direction = SOUTH
	case SOUTH:
		r.Direction = WEST
	case WEST:
		r.Direction = NORTH
	}
}

func (r *Rover) TurnLeft() {
	switch r.Direction {
	case NORTH:
		r.Direction = WEST
	case EAST:
		r.Direction = NORTH
	case SOUTH:
		r.Direction = EAST
	case WEST:
		r.Direction = SOUTH
	}
}
