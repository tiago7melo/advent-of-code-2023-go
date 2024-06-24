package main

import (
	"advent-of-code-2023-go/utils/fileutil"
	"fmt"
	"log"
	"strings"
	"unicode"
)

// Go doesn't allow complex types like slice and maps to be constant, so we'll just make a function to access the constant value instead
// this prevents mutation of the global values by other functions
func getDigitWords() []string {
	return []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
}

// Go doesn't allow complex types like slice and maps to be constant, so we'll just make a function to access the constant value instead
// this prevents mutation of the global values by other functions
func getWordToDigitMap() map[string]rune {
	wordToDigit := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	return wordToDigit
}

func findDigit(lb, rb, pos int, line string) (rune, bool) {
	digit := '0'
	found := false

	if unicode.IsDigit(rune(line[pos])) {
		digit = rune(line[pos])
		found = true
	} else {
		sl := line[lb:rb] // in line[low:high], low is included but high is excluded
		for _, dWord := range getDigitWords() {
			if strings.Contains(sl, dWord) {
				digit = getWordToDigitMap()[dWord]
				found = true
			}
		}
	}

	return digit, found
}

func getCalibrationVal(line string) int {
	var firstDigit, lastDigit rune = '0', '0'
	var firstDigitFound, lastDigitFound bool = false, false
	for firstDigitPos, _ := range line {

		if !firstDigitFound {
			firstDigit, firstDigitFound = findDigit(0, firstDigitPos+1, firstDigitPos, line)
		}

		if !lastDigitFound {
			lastDigitPos := len(line) - firstDigitPos - 1
			lastDigit, lastDigitFound = findDigit(lastDigitPos, len(line), lastDigitPos, line)
		}
	}
	return (int(firstDigit-'0') * 10) + int(lastDigit-'0')
}

func main() {
	var path string = "../input"
	lines, err := fileutil.ReadLines(path)
	if err != nil {
		log.Fatalf("Couldn't read lines of file %v", path)
		return
	}

	sum := 0
	for i, line := range lines {
		calibrationValue := getCalibrationVal(line)
		fmt.Printf("calVal for line %v: %v \n", i, calibrationValue)
		sum += calibrationValue
	}
	fmt.Printf("Sum: %v \n", sum)
}
