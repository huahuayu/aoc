package day3

type Position struct {
	X int
	Y int
}

func VisitedAtLeastOnce(route string) int {
	visited := map[Position]bool{}
	visited[Position{0, 0}] = true
	currentPosition := Position{0, 0}
	for _, direction := range route {
		currentPosition = Move(currentPosition, direction)
		visited[currentPosition] = true
	}

	return len(visited)
}

func Move(position Position, direction rune) Position {
	switch direction {
	case '^':
		return Position{position.X, position.Y + 1}
	case 'v':
		return Position{position.X, position.Y - 1}
	case '>':
		return Position{position.X + 1, position.Y}
	case '<':
		return Position{position.X - 1, position.Y}
	default:
		panic("unknown direction")
	}
}

func VisitedAtLeastOnceWithRoboSanta(route string) int {
	visited := map[Position]bool{}
	visited[Position{0, 0}] = true
	currentPosition := Position{0, 0}
	roboSantaPosition := Position{0, 0}
	for index, direction := range route {
		if index%2 == 0 {
			currentPosition = Move(currentPosition, direction)
			visited[currentPosition] = true
		} else {
			roboSantaPosition = Move(roboSantaPosition, direction)
			visited[roboSantaPosition] = true
		}
	}

	return len(visited)
}
