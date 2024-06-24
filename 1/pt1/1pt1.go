package main

import (
	"advent-of-code-2023-go/utils/fileutil"
	"fmt"
	"log"
	"unicode"
)

func main() {
	var path string = "../input"
	lines, err := fileutil.ReadLines(path)
	if err != nil {
		log.Fatalf("Couldn't read lines of file %v", path)
		return
	}

	sum := 0
	for i, s := range lines {
		var firstDigit, lastDigit rune = '0', '0'
		var firstDigitFound, lastDigitFound bool = false, false
		for firstDigitPos, _ := range s {

			if !firstDigitFound && unicode.IsDigit(rune(s[firstDigitPos])) {
				firstDigit = rune(s[firstDigitPos])
				firstDigitFound = true
			}

			if !lastDigitFound {
				lastDigitPos := len(s) - firstDigitPos - 1
				if unicode.IsDigit(rune(s[lastDigitPos])) {
					lastDigit = rune(s[lastDigitPos])
					lastDigitFound = true
				}
			}
		}

		calibrationValue := (int(firstDigit-'0') * 10) + int(lastDigit-'0')
		fmt.Printf("calVal for line %v: %v \n", i, calibrationValue)
		sum += calibrationValue
	}
	fmt.Printf("Sum: %v \n", sum)
}
