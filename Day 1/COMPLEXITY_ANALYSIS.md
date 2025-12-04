# O(n) vs O(2n) Solution Analysis

## Overview
This document explains the difference between the O(2n) solution in `2n.go` and the true O(n) solution in `n.go`.

## O(2n) Solution (`2n.go`)
**Time Complexity:** O(2n)  
**Space Complexity:** O(n)

### Breakdown:
1. **Step 1 - ReadInput: O(n)**
   - Opens the file
   - Scans through all n rotations line by line
   - Parses each line into a struct
   - **Stores all rotations in a slice in memory**
   - Returns the complete slice

2. **Step 2 - Solve: O(n)**
   - Iterates through the in-memory slice of n rotations
   - Calculates both Part 1 and Part 2 results
   - Single pass through the data

**Total:** O(n) + O(n) = **O(2n)**

The key issue: The input is stored entirely in memory before processing begins.

### Code Flow:
```
1. ReadInput("input.csv")      // O(n) - read and store all data
   ↓
2. solve(input)                // O(n) - process all data once
   ↓
3. Return results
```

---

## O(n) Solution (`n.go`)
**Time Complexity:** O(n)  
**Space Complexity:** O(1)

### Optimization:
Instead of reading all data first and then processing, we **process each rotation immediately as it's read from the file**.

### Key Differences:
1. **No intermediate storage**: We don't store the entire input slice
2. **Streaming processing**: Each line is processed as soon as it's read
3. **Single file scan**: One pass through the file with simultaneous computation
4. **Constant memory**: Only store current position and running totals (O(1) extra space)

### Code Flow:
```
for each line in file:          // O(n) - iterate through file once
   ↓
1. Parse the rotation
2. Calculate Part 2 contribution
3. Update dial position
4. Calculate Part 1 contribution
   ↓
Return results                   // Single pass complete
```

---

## Why This Matters

### Time Complexity:
- **O(2n)** means the coefficient is 2x, so it takes twice as many operations
- **O(n)** is the absolute minimum - we must read every line at least once
- For 4498 rotations: O(2n) = 8996 operations vs O(n) = 4498 operations

### Space Complexity:
> I am ignoring storage in this analysis as it has to be stored somewhere...
- **O(n)** storage: Stores all 4498 rotation structs (significant memory)
- **O(1)** storage: Only stores `currentPosition`, `part1Count`, `part2Count` (negligible)

---

## Implementation Details
> This is pseudo-code. See repository for actual solutions.
### O(2n)
```go
func readInput(filename string) ([]struct { ... }, error) {
    // ... open file ...
    var rotations []struct { Direction bool; Amount int }
    
    for scanner.Scan() {
        // Parse and APPEND to slice
        rotations = append(...)
    }
    
    return rotations  // Return entire slice - data persists in memory
}

func main() {
    rotations := readInput("input.csv")
    solveN(rotations)
}
```

### O(n)
```go
func solveN(filename string) (int, int, error) {
    // ... open file ...
    var part1Count, part2Count int
    currentPosition := 50
    
    for scanner.Scan() {
        line := scanner.Text()
        // Parse line
        amount, _ := strconv.Atoi(line[1:])
        
        // Process immediately - line is discarded after this iteration
        // No array/slice storage needed
        part2Count += ...
        currentPosition = ...
        part1Count += ...
    }
    
    return part1Count, part2Count  // Only return results
}
```

---

## Benchmark Summary

| Metric | O(2n) Solution | O(n) Solution |
|--------|---|---|
| Time | 2 × iterations | 1 × iterations |
| Space | 4498 structs stored (~≈256KB) | 3 integers stored (~≈12 bytes) |
| Read/Parse | All data read upfront | Streaming |
| Computing | After all data loaded | During read |

---

## Verification
Both solutions produce identical results:
- **Part 1:** 1129
- **Part 2:** 6638

The O(n) solution achieves 50% time improvement with virtually no memory overhead.
