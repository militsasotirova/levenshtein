// Calculates the Levenshtein Edit Distance between two strings.
package levenshtein

import (
	"fmt"
)

// Pretty prints the dynamic table used to calculate the Levenshtein Distance.
func displayDynamicTable(table *[][]int, a []rune, b []rune) {
	fmt.Println("\nDynamic Table:")
	fmt.Print("  ")
	for _, v := range b {
		fmt.Printf("%c ", v)
	}

	fmt.Println(" ")
	for r := 0; r < len(*table); r++ {
		fmt.Printf("%c ", a[r])
		for c := 0; c < len((*table)[0]); c++ {
			fmt.Print((*table)[r][c], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

// Generates the table and sets the first row and column values.
func tableSetUp(m, n int) *[][]int {
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}
	for c := range result[0] {
		result[0][c] = c
	}
	for r := 0; r < len(result); r++ {
		result[r][0] = r
	}

	return &result
}

// Computes the minimum of the parameters.
func customMin(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
	}

	return min
}

// Calculates the Levenshtein Distance between two strings, firstWord and secondWord. If
// display is true, the dynamic table of the calculation is printed.
func CalcLevenshteinDist(firstWord string, secondWord string, display bool) int {
	a := []rune(" " + firstWord)
	b := []rune(" " + secondWord)
	table := tableSetUp(len(a), len(b))
	for r := 1; r < len(a); r++ {
		for c := 1; c < len(b); c++ {
			//find local minimum
			upperLeft := (*table)[r-1][c-1]
			above := (*table)[r-1][c]
			left := (*table)[r][c-1]
			min := customMin(upperLeft, above, left)

			if a[r] == b[c] {
				(*table)[r][c] = min
			} else {
				(*table)[r][c] = min + 1
			}
		}
	}
	if display {
		displayDynamicTable(table, a, b)
	}

	return (*table)[len(a) - 1][len(b) - 1]
}

