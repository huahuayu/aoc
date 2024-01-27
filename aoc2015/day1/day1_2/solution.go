package day1_2

import "errors"

func floorCounter(instructions string, targetFloor int) (numOfSteps int, err error) {
	var floor int
	// Range the string to get the char
	for i, char := range instructions {
		if char == '(' {
			// Increment the floor counter
			floor++
		} else if char == ')' {
			// Decrement the floor counter
			floor--
		}
		if floor == targetFloor {
			return i + 1, nil
		}
	}
	return 0, errors.New("no solution")
}
