package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// solveN evaluates both parts in O(n) by processing rotations as they're read from the file
// Time Complexity: O(n) - single pass through input file, no intermediate storage
// Space Complexity: O(1) - only uses constant extra space (no storage of rotations slice)
func solveN(filename string) (int, int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	part1Count := 0
	part2Count := 0
	currentPosition := 50

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0] == 'R'
		amount, _ := strconv.Atoi(line[1:])

		// Calculate part 2: count passes through 0 during rotation
		if currentPosition == 0 {
			part2Count += amount / 100
		} else {
			distanceToZero := 100 - currentPosition
			if !direction {
				distanceToZero = currentPosition
			}
			if amount >= distanceToZero {
				part2Count += 1 + (amount-distanceToZero)/100
			}
		}

		// Update position
		if direction {
			currentPosition = (currentPosition + amount) % 100
		} else {
			currentPosition = ((currentPosition-amount)%100 + 100) % 100
		}

		// Calculate part 1: count landings on 0
		if currentPosition == 0 {
			part1Count++
		}
	}

	return part1Count, part2Count, scanner.Err()
}

func main() {
	part1Result, part2Result, err := solveN("../input.csv")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("=== Day 1: Secret Entrance (O(n) Solution) ===")
	fmt.Printf("Part 1: %d\n", part1Result)
	fmt.Printf("Part 2: %d\n", part2Result)
}
