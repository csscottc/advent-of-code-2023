package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {

	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var calibrationValues []int

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error reading line:", err)
			}
			break
		}
		calibrationValue := getCalibrationValue(line)
		num, err := strconv.Atoi(calibrationValue)

		if err != nil {
			fmt.Println(err)
		}
		calibrationValues = append(calibrationValues, num)
	}
	fmt.Println(sumCalibrationValues(calibrationValues))
}

func sumCalibrationValues(cvalues []int) int {
	var total = 0
	for _, val := range cvalues {
		total += val
	}
	return total
}

func getCalibrationValue(line string) string {
	var first rune
	var last rune
	var firstAssigned = false
	for _, rv := range line {
		if unicode.IsDigit(rune(rv)) {
			if firstAssigned == false {
				first = rv
				firstAssigned = true
			}
			last = rv
		}
	}
	return fmt.Sprintf("%c%c", first, last)
}
