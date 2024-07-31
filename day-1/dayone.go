package dayone

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

/*
The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover.
On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

What is the sum of all of the calibration values?

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.
*/

func getNumberFromString(s string) int {
	var firstDigit int
	var lastDigit int

	r := []rune(s) // r - runes

	// get last digit
	for i := len(r) - 1; i >= 0; i-- {
		if unicode.IsDigit(r[i]) {
			lastDigit = int(r[i] - '0') // r[i] is unicode/ASCII not the int itself - subtract 0's ASCII from it.
			break
		}
	}

	// get first digit
	for i := 0; i < len(r); i++ {
		if unicode.IsDigit(r[i]) {
			firstDigit = int(r[i] - '0')
			break
		}
	}

	return 10*firstDigit + lastDigit
}

func CalibrationSum() int {

	// read file
	file, err := os.Open("day-1/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += getNumberFromString(line)
	}

	return sum

}
