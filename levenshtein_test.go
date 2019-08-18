package levenshtein

import (
	"testing"
)

// Test the case where both strings are empty.
func TestTwoEmpty(t *testing.T) {
	firstWord := ""
	secondWord := ""
	editDistance := CalcLevenshteinDist(firstWord, secondWord, false)
	expected := 0
	if editDistance != expected {
		t.Errorf("Edit distance was incorrect, got: %d, want: %d.", editDistance, expected)
	}
}

// Tests the case where the first argument is longer than the second.
func TestFirstLonger(t *testing.T) {
	firstWord := "hyundai"
	secondWord := "honda"
	editDistance := CalcLevenshteinDist(firstWord, secondWord, false)
	expected := 3
	if editDistance != expected {
		t.Errorf("Edit distance was incorrect, got: %d, want: %d.", editDistance, expected)
	}
}

// Tests the case where the second argument is longer than the first.
func TestSecondLonger(t *testing.T) {
	firstWord := "honda"
	secondWord := "hyundai"
	editDistance := CalcLevenshteinDist(firstWord, secondWord, false)
	expected := 3
	if editDistance != expected {
		t.Errorf("Edit distance was incorrect, got: %d, want: %d.", editDistance, expected)
	}
}

// Tests the case where the both arguments are of equal length.
func TestEqualLength(t *testing.T) {
	firstWord := "hello"
	secondWord := "world"
	editDistance := CalcLevenshteinDist(firstWord, secondWord, false)
	expected := 4
	if editDistance != expected {
		t.Errorf("Edit distance was incorrect, got: %d, want: %d.", editDistance, expected)
	}
}
