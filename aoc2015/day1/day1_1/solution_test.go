package day1_1

import (
	"aoc/aoc2015/day1/input"
	"testing"
)

func TestSolution(t *testing.T) {
	result := floorCounter(input.Instructions)
	t.Log(result)
}
