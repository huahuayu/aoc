package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type Box struct {
	Length int
	Width  int
	Height int
}

func NewBox(boxSpec string) (*Box, error) {
	specs := strings.Split(boxSpec, "x")
	if len(specs) != 3 {
		return nil, fmt.Errorf("invalid box spec: %s", boxSpec)
	}
	// Convert the specs to int
	l, err := strconv.Atoi(specs[0])
	if err != nil {
		return nil, fmt.Errorf("invalid box spec length: %s", boxSpec)
	}
	w, err := strconv.Atoi(specs[1])
	if err != nil {
		return nil, fmt.Errorf("invalid box spec width: %s", boxSpec)
	}
	h, err := strconv.Atoi(specs[2])
	if err != nil {
		return nil, fmt.Errorf("invalid box spec height: %s", boxSpec)
	}
	return &Box{
		Length: l,
		Width:  w,
		Height: h,
	}, nil
}

func (box *Box) WrappingPaperArea() int {
	return box.surfaceArea() + box.smallestSideArea()
}

func (box *Box) RibbonLength() int {
	return box.smallestPerimeter() + box.cubicVolume()
}

func (box *Box) surfaceArea() int {
	return 2*box.Length*box.Width + 2*box.Width*box.Height + 2*box.Height*box.Length
}

func (box *Box) smallestSideArea() int {
	return min(box.Length*box.Width, box.Width*box.Height, box.Height*box.Length)
}

func (box *Box) cubicVolume() int {
	return box.Length * box.Width * box.Height
}

func (box *Box) smallestPerimeter() int {
	return min(2*(box.Length+box.Width), 2*(box.Width+box.Height), 2*(box.Height+box.Length))
}
