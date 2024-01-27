package day3

import (
	"aoc/aoc2015/day3/input"
	"testing"
)

func TestVisitedAtLeastOnce(t *testing.T) {
	t.Log(VisitedAtLeastOnce(input.Route))
	t.Log(VisitedAtLeastOnceWithRoboSanta(input.Route))
}
