package day6

import (
	"regexp"
	"strconv"
	"strings"
)

type Grid struct {
	Cells      [][]bool
	Brightness [][]int
}

type Point struct {
	X int
	Y int
}

type CoordinatePair struct {
	Start Point
	End   Point
}

type Action string

const (
	Unknown Action = ""
	TurnOn  Action = "turn on"
	TurnOff Action = "turn off"
	Toggle  Action = "toggle"
)

var re = regexp.MustCompile(`\d+`)

func NewGrid() *Grid {
	cells := make([][]bool, 1000)
	brightness := make([][]int, 1000)
	for i := range cells {
		cells[i] = make([]bool, 1000)
		brightness[i] = make([]int, 1000)
	}
	return &Grid{Cells: cells, Brightness: brightness}
}

func (grid *Grid) TurnOn(coordinatePair CoordinatePair) {
	for x := coordinatePair.Start.X; x <= coordinatePair.End.X; x++ {
		for y := coordinatePair.Start.Y; y <= coordinatePair.End.Y; y++ {
			grid.Cells[x][y] = true
			grid.Brightness[x][y]++
		}
	}
}

func (grid *Grid) TurnOff(coordinatePair CoordinatePair) {
	for x := coordinatePair.Start.X; x <= coordinatePair.End.X; x++ {
		for y := coordinatePair.Start.Y; y <= coordinatePair.End.Y; y++ {
			grid.Cells[x][y] = false
			if grid.Brightness[x][y] > 0 {
				grid.Brightness[x][y]--
			}
		}
	}
}

func (grid *Grid) Toggle(coordinatePair CoordinatePair) {
	for x := coordinatePair.Start.X; x <= coordinatePair.End.X; x++ {
		for y := coordinatePair.Start.Y; y <= coordinatePair.End.Y; y++ {
			grid.Cells[x][y] = !grid.Cells[x][y]
			grid.Brightness[x][y] += 2
		}
	}
}

func (grid *Grid) FollowInstruction(instruction string) {
	action, coordinatePair := parseInstruction(instruction)
	if action == TurnOn {
		grid.TurnOn(coordinatePair)
		return
	}
	if action == TurnOff {
		grid.TurnOff(coordinatePair)
		return
	}
	if action == Toggle {
		grid.Toggle(coordinatePair)
	}

}

func (grid *Grid) CountLightsOn() int {
	var count int
	for x := range grid.Cells {
		for y := range grid.Cells[x] {
			if grid.Cells[x][y] {
				count++
			}
		}
	}
	return count
}

func (grid *Grid) TotalBrightness() int {
	var count int
	for x := range grid.Brightness {
		for y := range grid.Brightness[x] {
			count += grid.Brightness[x][y]
		}
	}
	return count
}

func parseInstruction(instruction string) (Action, CoordinatePair) {
	var action Action
	if strings.HasPrefix(instruction, string(TurnOn)) {
		action = TurnOn
	} else if strings.HasPrefix(instruction, string(TurnOff)) {
		action = TurnOff
	} else if strings.HasPrefix(instruction, string(Toggle)) {
		action = Toggle
	} else {
		return Unknown, CoordinatePair{}
	}

	// Find all matches in the string
	matches := re.FindAllString(instruction, -1)

	// Convert the resulting strings to integers
	numbers := make([]int, len(matches))
	for i, match := range matches {
		number, _ := strconv.Atoi(match)
		numbers[i] = number
	}

	// Validate the points
	if len(numbers) != 4 {
		return Unknown, CoordinatePair{}
	}

	coordinatePair := CoordinatePair{
		Start: Point{
			X: numbers[0],
			Y: numbers[1],
		},
		End: Point{
			X: numbers[2],
			Y: numbers[3],
		},
	}
	return action, coordinatePair
}
