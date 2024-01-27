package day4

import (
	"aoc/aoc2015/day4/input"
	"fmt"
	"testing"
)

func TestMD5SaltFinder(t *testing.T) {
	t.Log(MD5SaltFinder(input.Input, 6))
}

func TestMD5(t *testing.T) {
	t.Log(MD5(fmt.Sprintf("%s%d", input.Input, 254575)))
}
