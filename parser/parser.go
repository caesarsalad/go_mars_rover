package parser

import (
	"errors"
	"strconv"
	"strings"

	"github.com/caesarsalad/go_mars_rover/rover"
)

var (
	validInstractions = []string{rover.MOVE, rover.LEFT, rover.RIGHT}
	validcoordinates  = []string{rover.NORTH, rover.EAST, rover.SOUTH, rover.WEST}
)

func coordinateParser(x_input, y_input string) (int, int, error) {
	var x, y int

	x, err := strconv.Atoi(x_input)
	if err != nil {
		return x, y, errors.New("x coordinate must be int")
	}
	if x <= 0 {
		return x, y, errors.New("x coordinate must be greater than zero")
	}

	y, err = strconv.Atoi(y_input)
	if err != nil {
		return x, y, errors.New("y coordinate must be int")
	}
	if y <= 0 {
		return x, y, errors.New("y coordinate must be greater than zero")
	}
	return x, y, nil
}

func ParseValidatePlateau(input string) (int, int, error) {
	input_list := strings.Split(input, " ")
	if len(input_list) != 2 {
		return 0, 0, errors.New("plateau coordinates must be X and Y format with space")
	}
	x, y, err := coordinateParser(input_list[0], input_list[1])
	if err != nil {
		return x, y, err
	}
	return x, y, nil
}

func ParseValidateRoverPosition(input string) (int, int, string, error) {
	input_list := strings.Split(input, " ")
	if len(input_list) != 3 {
		return 0, 0, "", errors.New("rover coordinates must be X, Y and rover front direction format with space")
	}

	x, y, err := coordinateParser(input_list[0], input_list[1])
	if err != nil {
		return x, y, "", err
	}

	direction := input_list[2]
	if !contains(validcoordinates, direction) {
		return x, y, "", errors.New("invalid direction for the rover")
	}

	return x, y, direction, nil
}

func ParseValidateRoverInstructions(input string) ([]string, error) {
	input_list := strings.Split(input, "")
	if len(input_list) == 0 {
		return input_list, errors.New("instructions are empty")
	}
	for _, instraction := range input_list {
		if !contains(validInstractions, instraction) {
			return input_list, errors.New("invalid instraction")
		}
	}
	return input_list, nil
}
