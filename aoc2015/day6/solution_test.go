package day6

import (
	"aoc/aoc2015/day6/input"
	"fmt"
	"testing"
)

// TestBackTest tests by DemoInput
func TestBackTest(t *testing.T) {
	grid := NewGrid()
	for _, instruction := range input.DemoInput {
		grid.FollowInstruction(instruction)
	}
	// Expected 998,996
	if grid.CountLightsOn() != 998996 {
		t.Errorf("Expected 998996, got %d", grid.CountLightsOn())
	}
}

func TestDay6Part1(t *testing.T) {
	grid := NewGrid()
	for _, instruction := range input.Input {
		grid.FollowInstruction(instruction)
	}
	fmt.Println(grid.CountLightsOn(), grid.TotalBrightness())
}
