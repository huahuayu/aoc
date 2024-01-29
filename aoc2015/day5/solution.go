package day5

import "strings"

type StringType string

const (
	Nice    StringType = "Nice"
	Naughty StringType = "Naughty"
)

var forbiddenList = []string{"ab", "cd", "pq", "xy"}

func StringClassifier(stringToCheck string, niceRules []func(string) bool) StringType {
	for _, rule := range niceRules {
		if !rule(stringToCheck) {
			return Naughty
		}
	}
	return Nice
}

func atLeastThreeVowels(stringToCheck string) bool {
	vowels := 0
	for _, char := range stringToCheck {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}
	}
	return vowels >= 3
}

func atLeastOneLetterTwiceInARow(stringToCheck string) bool {
	for i := 1; i < len(stringToCheck); i++ {
		if stringToCheck[i-1] == stringToCheck[i] {
			return true
		}
	}
	return false
}

func forbiddenStrings(stringToCheck string) bool {
	for _, forbiddenString := range forbiddenList {
		if strings.Contains(stringToCheck, forbiddenString) {
			return false
		}
	}
	return true
}

func part2RulePairOfTwoLettersAppearsTwice(stringToCheck string) bool {
	for i := 2; i < len(stringToCheck); i++ {
		if strings.Contains(stringToCheck[i:], stringToCheck[i-2:i]) {
			return true
		}
	}
	return false
}

func part2RuleLetterRepeatsWithOneLetterBetween(stringToCheck string) bool {
	for i := 2; i < len(stringToCheck); i++ {
		if stringToCheck[i-2] == stringToCheck[i] {
			return true
		}
	}
	return false
}
