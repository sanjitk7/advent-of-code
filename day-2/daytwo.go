package daytwo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func compareHashMaps(d map[string]int) bool {
	if d["R"] <= truthMap["R"] && d["G"] <= truthMap["G"] && d["B"] <= truthMap["B"] {
		return true
	}

	return false
}

// check if this game is possible
func processGame(s string) bool {
	parts := strings.Split(s, ":")

	if len(parts) < 2 {
		return false
	}

	// extract game number
	gameParts := strings.Split(parts[0], " ")
	gameNum, err := strconv.Atoi(gameParts[1])

	if err != nil {
		fmt.Println("Panic! incorrect gameNum parsed!")
	}

	// extract each draw
	drawParts := strings.Split(parts[1], ";")

	for _, draw := range drawParts {
		// create a map for each draw
		curDrawMap := make(map[string]int)
		colorParts := strings.Split(draw, ",")

		for _, color := range colorParts {
			colorNumParts := strings.Split(strings.TrimSpace(color), " ")
			// fmt.Println("colorNumParts: ", colorNumParts[0], len(colorNumParts))
			colorNumber, err := strconv.Atoi(colorNumParts[0])

			if err != nil {
				fmt.Println("strconv error: ", err)
			}
			if colorNumParts[1] == "red" {
				curDrawMap["R"] = colorNumber
			} else if colorNumParts[1] == "blue" {
				curDrawMap["B"] = colorNumber
			} else if colorNumParts[1] == "green" {
				curDrawMap["G"] = colorNumber
			} else {
				fmt.Println("color found in draw not R, B, G!")
			}
		}
		// check if this draw is possible
		if !compareHashMaps(curDrawMap) {
			return false
		}
	}
	fmt.Println("Possible gameNum: %d", gameNum)
	PossibleGameSum += gameNum
	return true

}

var truthMap = make(map[string]int)
var PossibleGameSum int

func CubeConundrum() int {

	truthMap["R"] = 12
	truthMap["G"] = 13
	truthMap["B"] = 14

	PossibleGameSum = 0

	// extract game map
	file, err := os.Open("day-2/input.txt")
	if err != nil {
		fmt.Println("error opening file!\n")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// var gameMap map[string]int
	for scanner.Scan() {
		line := scanner.Text()
		processGame(line)
	}

	return PossibleGameSum
}
