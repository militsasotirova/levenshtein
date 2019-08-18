Calculate the Levenshtein Edit Distance between two strings.

# Install
```
$ go get github.com/militsasotirova/levenshtein
```
# Use
```
package main

import (
	"bufio"
	"fmt"
	"github.com/militsasotirova/levenshtein"
	"os"
	"strings"
)

func main() {
	fmt.Println("string to string correction")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("First word: ")
	firstWord, _ := reader.ReadString('\n')
	firstWord = strings.TrimSpace(firstWord)
	fmt.Print("Second word: ")
	secondWord, _ := reader.ReadString('\n')
	secondWord = strings.TrimSpace(secondWord)
	editDistance := levenshtein.CalcLevenshteinDist(firstWord, secondWord, true)
	fmt.Printf("The edit distance (Levenshtein Distance) between %s and %s is %d.\n", firstWord, secondWord, editDistance)
}

```

