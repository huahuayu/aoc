package day1_1

func floorCounter(instructions string) int {
	var floor int
	// Range the string to get the char
	for _, char := range instructions {
		if char == '(' {
			// Increment the floor counter
			floor++
		} else if char == ')' {
			// Decrement the floor counter
			floor--
		}
	}
	return floor
}
