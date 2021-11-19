package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/caesarsalad/go_mars_rover/parser"
	"github.com/caesarsalad/go_mars_rover/rover"
)

func create_plateu() (rover.Plateau, error) {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter plateau area bound coordinates: ")
	scanner.Scan()
	input = scanner.Text()
	x, y, err := parser.ParseValidatePlateau(input)
	if err != nil {
		return rover.Plateau{}, err
	}
	plateau := rover.Plateau{X_bound: x, Y_bound: y}

	return plateau, nil
}

func go_rover(plateau rover.Plateau) {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter rover position: ")
	scanner.Scan()
	input = scanner.Text()
	var direction string
	x, y, direction, err := parser.ParseValidateRoverPosition(input)
	if err != nil {
		log.Println(err)
		return
	}
	position := rover.Position{X: x, Y: y}
	mars_rover := rover.Rover{Position: &position,
		Plateau:   plateau,
		Direction: direction}

	fmt.Print("Enter rover instruction: ")
	scanner.Scan()
	input = scanner.Text()
	var instructions []string
	instructions, err = parser.ParseValidateRoverInstructions(input)
	if err != nil {
		log.Println(err)
		return
	}
	err = mars_rover.ReadInstructions(instructions)

	if err != nil {
		log.Println(err)
	}

	fmt.Println("rover stoped at ", mars_rover.Position.X, mars_rover.Position.Y, mars_rover.Direction)
}

func main() {
	plateu, err := create_plateu()
	if err != nil {
		log.Fatalln(err)
	}
	go_rover(plateu)

	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Do you want another rover? y/n ")
		scanner.Scan()
		input = scanner.Text()
		if input == "y" {
			go_rover(plateu)
		} else {
			log.Println("Exiting...")
			os.Exit(0)
		}
	}
}
