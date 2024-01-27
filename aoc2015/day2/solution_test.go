package day2

import (
	"aoc/aoc2015/day2/input"
	"testing"
)

func TestBox_WrappingPaperArea(t *testing.T) {
	var totalArea int
	var totalRibbonLength int
	for _, boxSpec := range input.BoxSpecs {
		box, err := NewBox(boxSpec)
		if err != nil {
			t.Fatal(err)
		}
		totalArea += box.WrappingPaperArea()
		totalRibbonLength += box.RibbonLength()
	}
	t.Log(totalArea)
	t.Log(totalRibbonLength)
}
