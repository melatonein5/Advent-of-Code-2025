/*
--- Day 1: Secret Entrance ---
You arrive at the secret entrance to the North Pole base ready to start decorating. Unfortunately, the password seems to have been changed, so you can't get in. A document taped to the wall helpfully explains:

"Due to new security protocols, the password is locked in the safe below. Please see the attached document for the new combination."

The safe has a dial with only an arrow on it; around the dial are the numbers 0 through 99 in order. As you turn the dial, it makes a small click noise as it reaches each number.

The attached document (your puzzle input) contains a sequence of rotations, one per line, which tell you how to open the safe. A rotation starts with an L or R which indicates whether the rotation should be to the left (toward lower numbers) or to the right (toward higher numbers). Then, the rotation has a distance value which indicates how many clicks the dial should be rotated in that direction.

So, if the dial were pointing at 11, a rotation of R8 would cause the dial to point at 19. After that, a rotation of L19 would cause it to point at 0.

Because the dial is a circle, turning the dial left from 0 one click makes it point at 99. Similarly, turning the dial right from 99 one click makes it point at 0.

So, if the dial were pointing at 5, a rotation of L10 would cause it to point at 95. After that, a rotation of R5 could cause it to point at 0.

The dial starts by pointing at 50.

You could follow the instructions, but your recent required official North Pole secret entrance security training seminar taught you that the safe is actually a decoy. The actual password is the number of times the dial is left pointing at 0 after any rotation in the sequence.

For example, suppose the attached document contained the following rotations:

L68
L30
R48
L5
R60
L55
L1
L99
R14
L82

Following these rotations would cause the dial to move as follows:

    The dial starts by pointing at 50.
    The dial is rotated L68 to point at 82.
    The dial is rotated L30 to point at 52.
    The dial is rotated R48 to point at 0.
    The dial is rotated L5 to point at 95.
    The dial is rotated R60 to point at 55.
    The dial is rotated L55 to point at 0.
    The dial is rotated L1 to point at 99.
    The dial is rotated L99 to point at 0.
    The dial is rotated R14 to point at 14.
    The dial is rotated L82 to point at 32.

Because the dial points at 0 a total of three times during this process, the password in this example is 3.

Analyze the rotations in your attached document. What's the actual password to open the door?
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

// Right turns adds to the current direction
func turnRight(current, amount int) int {
	//Firstly, perform the addition
	newValue := current + amount

	//Next, as the number needs to be in the range of 0 to 99, we can use modulo
	newValue = newValue % 100
	return newValue
}

// Left turns subtracts from the current direction
func turnLeft(current, amount int) int {
	//Firstly, perform the subtraction
	newValue := current - amount

	//For modulo with negative numbers on a circular dial, use this formula
	newValue = ((newValue % 100) + 100) % 100
	return newValue
}

// ReadInput reads the input file and returns a cuple, one boolean indicating the turn direction (true for right, false for left) and an integer indicating the amount to turn
func ReadInput(filename string) ([]struct {
	Direction bool
	Amount    int
}, error) {
	//Open the file
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	//Create a scanner
	scanner := bufio.NewScanner(f)

	//Each line will start with L or R, followed by a number
	var rotations []struct {
		Direction bool
		Amount    int
	}

	for scanner.Scan() {
		line := scanner.Text()
		var direction bool
		if line[0] == 'R' {
			direction = true
		} else {
			direction = false
		}

		amount := line[1:] //Get the substring after the first character
		var amountInt int
		fmt.Sscanf(amount, "%d", &amountInt)

		rotations = append(rotations, struct {
			Direction bool
			Amount    int
		}{Direction: direction, Amount: amountInt})
	}

	return rotations, nil
}

func main() {
	//Create a counter for the number of times we land on 0
	zeroCount := 0

	//Read the input
	rotations, err := ReadInput("part1.csv")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	//The dial starts at 50
	currentPosition := 50

	//Process each rotation
	for _, rotation := range rotations {
		if rotation.Direction {
			//Right turn
			currentPosition = turnRight(currentPosition, rotation.Amount)
		} else {
			//Left turn
			currentPosition = turnLeft(currentPosition, rotation.Amount)
		}

		//Check if we are at 0
		if currentPosition == 0 {
			zeroCount++
		}
	}

	//Output the result
	fmt.Println("The password to open the door is:", zeroCount)
}
