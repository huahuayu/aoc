package day5

import (
	"aoc/aoc2015/day5/input"
	"testing"
)

func TestStringClassifierBackTest(t *testing.T) {
	for _, testCase := range []struct {
		input    string
		expected StringType
	}{
		{"ugknbfddgicrmopn", Nice},
		{"aaa", Nice},
		{"jchzalrnumimnmhp", Naughty},
		{"haegwjzuvuyypxyu", Naughty},
		{"dvszwmarrgswjxmb", Naughty},
	} {
		result := StringClassifier(testCase.input, []func(string) bool{atLeastThreeVowels, atLeastOneLetterTwiceInARow, forbiddenStrings})
		if result != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, result)
		}
	}
}

func TestStringClassifier(t *testing.T) {
	niceCount := 0
	for _, str := range input.Input {
		if StringClassifier(str, []func(string) bool{atLeastThreeVowels, atLeastOneLetterTwiceInARow, forbiddenStrings}) == Nice {
			niceCount++
		}
	}
	t.Log(niceCount)
}

func TestPart2StringClassifierBackTest(t *testing.T) {
	for _, testCase := range []struct {
		input    string
		expected StringType
	}{
		{"qjhvhtzxzqqjkmpb", Nice},
		{"xxyxx", Nice},
		{"uurcxstgmygtbstg", Naughty},
		{"ieodomkazucvgmuy", Naughty},
	} {
		result := StringClassifier(testCase.input, []func(string) bool{part2RulePairOfTwoLettersAppearsTwice, part2RuleLetterRepeatsWithOneLetterBetween})
		if result != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, result)
		}
	}
}

func TestPart2StringClassifier(t *testing.T) {
	niceCount := 0
	for _, str := range input.Input {
		if StringClassifier(str, []func(string) bool{part2RulePairOfTwoLettersAppearsTwice, part2RuleLetterRepeatsWithOneLetterBetween}) == Nice {
			niceCount++
		}
	}
	t.Log(niceCount)
}
