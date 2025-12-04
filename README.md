# Advent of Code 2025
> Author: Joseph McCallum-Nattrass, Advent of Code.

Advent of Code is an Advent calendar of small programming puzzles for a variety of skill levels that can be solved in any programming language you like. People use them as interview prep, company training, university coursework, practice problems, a speed contest, or to challenge each other.

> Any text that appears like the following was written by me, while the rest is copied and pasted from the puzzle. Any code that appears was written by me, unless stated otherwise (often in sections labelled "Other Oddities"). Note that this only applies to this file. Any other files are authored by myself.

> All code is written in Go, as that is my language of preference due to its strong standard library (although much of it will go unused due to how low-level the puzzles are).

> Each users is generated a random set of inputs to use, so my answers will be unique to myself.

## Day 1: Secret Entrance
> There are two parts to this problem.

This problem involves a safe with a dial. The safe has a dial with only an arrow on it; around the dial are the numbers 0 through 99 in order. As you turn the dial, it makes a small click noise as it reaches each number.

The puzzle input contains a sequence of rotations, one per line, which tell you how to open the safe. A rotation starts with an L or R which indicates whether the rotation should be to the left (toward lower numbers) or to the right (toward higher numbers). Then, the rotation has a distance value which indicates how many clicks the dial should be rotated in that direction.

So, if the dial were pointing at 11, a rotation of R8 would cause the dial to point at 19. After that, a rotation of L19 would cause it to point at 0.

Because the dial is a circle, turning the dial left from 0 one click makes it point at 99. Similarly, turning the dial right from 99 one click makes it point at 0.

So, if the dial were pointing at 5, a rotation of L10 would cause it to point at 95. After that, a rotation of R5 could cause it to point at 0.

The dial starts by pointing at 50.

> **Time Complexity** : O(n) & O(2n)
> **Space Complexity**: O(1) & O(n) 

> Space complexity of the completed puzzle is always technically going to be O(n), the input still needs to be stored on disk, which counts as auxiliary space, contributing to the space complexity. It will naturally grow in a linear fashion. While my code has been modified to read from disk as a data stream and evaluates the rotation as it is being read there and then, 

> My first attempt resulted in a time complexity of O(3n). This is broken down into: reading the file into a data structure in memory; solving the part 1; solving part 2. This is because I was reading the input into a data structure, and then performing my algorithms separately. I then condensed this down into O(2n) complexity by combining the two algorithms to evaluate alongside each other. This was then further condensed into a total time complexity of O(n) by evaluating a rotation when it is read from the input file.

### Part 1: Counting how many times the an input results in the dial pointing at 0
For this puzzle, the solution is how many times the arrow on the dial points to 0.

For example, suppose the input contained the following rotations:
```
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
```

Following these rotations would cause the dial to move as follows:
```
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
```

Because the dial points at 0 a total of three times during this process, the solution in this example is 3.

>The first thing I noticed is that there are some big numbers in the input. Input 128 is a right turn by 958. For my solution, I wrote two functions. One to evaluate a left turn and another to evaluate a right turn:
```go
// turnRight rotates the dial to the right by the specified amount
func turnRight(current, amount int) int {
	return (current + amount) % 100
}

// turnLeft rotates the dial to the left by the specified amount
func turnLeft(current, amount int) int {
	return ((current-amount) % 100 + 100) % 100
}
```
> For the right turn, it simply adds the numbers together and performs a modulo function (`% 100`), which divides the sum by 100 and returns the remainder.

In computing and mathematics, the modulo operation returns the remainder or signed remainder of a division, after one number is divided by another, the latter being called the modulus of the operation (Wikipedia).
> This is performed because the dial is circular. If the number is bigger than 99, it needs to normalised to a number between 0 & 99. In decimal systems, you would simply just look at the last two digits. As the data is stored in binary, a mathematical operation is required. I could probably convert to a string and get the last two characters at the cost of memory usage, but it doesn't look as elegant. I also haven't benchmarked the two approaches. While division and multiplication operations are often computationally expensive, I imagine the amount on conversion going on would benefit performance, and may actually hinder it. For context, this is what that would look like:
```go
// String Manipulation (Not Recommended - For Reference Only)
func turnRightString(current, amount int) int {
    sum := current + amount
    // Convert to string, get last 2 digits, convert back to int
    sumStr := strconv.Itoa(sum)
    if len(sumStr) > 2 {
        sumStr = sumStr[len(sumStr)-2:]
    }
    result, _ := strconv.Atoi(sumStr)
    return result
}

func turnLeftString(current, amount int) int {
    diff := current - amount
    // Handle negative numbers by adding 100 until positive
    for diff < 0 {
        diff += 100
    }
    diffStr := strconv.Itoa(diff)
    if len(diffStr) > 2 {
        diffStr = diffStr[len(diffStr)-2:]
    }
    result, _ := strconv.Atoi(diffStr)
    return result
}
```
> The left turn is similar to the right turn, but I subtract instead of add. The difference here is that subtraction can lead to negative numbers outside the range of 0 to 99, so a modulo operations is performed to lower the range to -99 to 98 (it is impossible to get 99, as that is the maximum value, I am subtracting, and 0 is not a valid input for the magnitude of the turn). I then add 100, which is a computationally cheap way of ensuring the number is positive (the minimum value is -99, which results in a new value of 1). If the number is positive, it will now be above 100 (the maximum value possible here is 198), so another modulo function is required ensure it is within the range of 0 & 99.

> I can now check if the number is 0. If it is 0, I can just increment a counter.
```go
// Calculate part 1: count landings on 0
		if currentPosition == 0 {
			part1Count++
		}
```
> Once all the inputs have been evaluated, I simply print the value of `part1Count`, which is 1129 for the inputs I was provided.

### Part 2: Counting how many times the dial has a value of 0
I now have to count the number of times any click causes the dial to point at 0, regardless of whether it happens during a rotation or at the end of one.

Following the same rotations as in the above example, the dial points at zero a few extra times during its rotations:
```
The dial starts by pointing at 50.
The dial is rotated L68 to point at 82; during this rotation, it points at 0 once.
The dial is rotated L30 to point at 52.
The dial is rotated R48 to point at 0.
The dial is rotated L5 to point at 95.
The dial is rotated R60 to point at 55; during this rotation, it points at 0 once.
The dial is rotated L55 to point at 0.
The dial is rotated L1 to point at 99.
The dial is rotated L99 to point at 0.
The dial is rotated R14 to point at 14.
The dial is rotated L82 to point at 32; during this rotation, it points at 0 once.
```
In this example, the dial points at 0 three times at the end of a rotation, plus three more times during a rotation. So, in this example, the new solution would be 6.

If the dial were pointing at 50, a single rotation like R1000 would cause the dial to point at 0 ten times before returning back to 50!

> For this problem, I came up with 3 principles:
> 1. **For a right turn**: If the rotation amount is greater than or equal to (100 - currentPosition), the dial will pass through 0.
> 2. **For a left turn**: If the rotation amount is greater than or equal to currentPosition, the dial will pass through 0.
> 3. **Multiple passes**: If the rotation amount is large enough, the dial can pass through 0 multiple times. The number of additional passes after the first is calculated by `(rotation.Amount - distanceToZero) / 100`.

> Using these principles, I can calculate how many times the dial passes through 0 during each rotation:

```go
// Calculate part 2: count passes through 0 during rotation
if currentPosition == 0 {
    part2Count += rotation.Amount / 100
} else {
    distanceToZero := 100 - currentPosition
    if !rotation.Direction {
        distanceToZero = currentPosition
    }
    if rotation.Amount >= distanceToZero {
        part2Count += 1 + (rotation.Amount-distanceToZero)/100
    }
}

// Update position
if rotation.Direction {
    currentPosition = turnRight(currentPosition, rotation.Amount)
} else {
    currentPosition = turnLeft(currentPosition, rotation.Amount)
}
```

> If the current position is already at 0, then I calculate how many complete rotations of 100 occur during this turn using integer division (`rotation.Amount / 100`). If the current position is not 0, I calculate the distance to 0. For a right turn, this is `100 - currentPosition`. For a left turn, this is simply `currentPosition`. If the rotation amount is at least this distance, then we will pass through 0 at least once. The total count is `1` (for the first pass) plus any additional complete rotations: `(rotation.Amount - distanceToZero) / 100`.

> Once all the inputs have been evaluated, I simply print the value of `part2Count`, which is 6638 for the inputs I was provided.

#### Other Oddities
This section is my own work. There are two other things worth mentioning:
- How I read the input.
- How I parse the input.

I simply copied and pasted my input file into a comma separated values (`.csv`) file, and use the Go standard library to real the file line-by-line. Each rotation is stored on it's own line.
```go
// readInput parses the input file and returns a slice of rotations
func readInput(filename string) ([]struct {
	Direction bool
	Amount    int
}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var rotations []struct {
		Direction bool
		Amount    int
	}

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0] == 'R'
		var amount int
		fmt.Sscanf(line[1:], "%d", &amount)

		rotations = append(rotations, struct {
			Direction bool
			Amount    int
		}{Direction: direction, Amount: amount})
	}

	return rotations, nil
}
```

The first character will always be the direction of the rotation (`R` or `L`). I am storing this as a boolean, with `R` being represented by `true` and `L`  being representing by `false` (in fact, anything that isn't an `R` is represented by `false`, which leads to behaviour where `☭345` is a left turn by 345, and I supposed it is appropriate that a hammer and sickle glyph results in a rotation to the left).

A boolean takes up a single byte of memory. This is as memory efficient as storing `R` or `L`, which also can be represent as a single byte, being represented by `0x52` and `0x4C` respectively. Using a boolean means I can have more elegant code for statements logical comparators. See below for an example:
```go
// Storing R and L as bytes (BAD!!!)
if rotation.Direction == 'R' {
    currentPosition = turnRight(currentPosition, rotation.Amount)   
} else if rotation.Direction == 'L' {
    currentPosition = turnLeft(currentPosition, rotation.Amount)
}

// Storing R and L as a boolean (GOOD)
if rotation.Direction {
    currentPosition = turnRight(currentPosition, rotation.Amount)   
} else {
    currentPosition = turnLeft(currentPosition, rotation.Amount)
}
```
The boolean approach is cleaner, more readable, and easier to reason about. Additionally, when calculating the distance to 0 during Part 2, I can use the boolean directly in a logical expression:
```go
distanceToZero := 100 - currentPosition
if !rotation.Direction {
    distanceToZero = currentPosition
}
```
This is far more elegant than comparing characters or using switch statements.

For how much you actually turn the dial, we know that the first character of each line can be ignored. We can just chop it off and store the number with `fmt.Sscanf(line[1:], "%d", &amount)`. Easy. 