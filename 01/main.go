package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
  "strings"
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
		calibrationValue := GetCalibrationValue(line)
		calibrationValues = append(calibrationValues, calibrationValue)
	}
  fmt.Println("The sum is:",SumCalibrationValues(calibrationValues))
}

func lookup(word string) (rune, bool) {
  m := map[string]rune{
    "one": '1',
    "two": '2',
    "three": '3',
    "four": '4',
    "five": '5',
    "six": '6',
    "seven": '7',
    "eight": '8',
    "nine": '9',
  }
  val, ok := m[word]
  return val, ok
}

func SumCalibrationValues(cvalues []int) int {
	var total = 0
	for _, val := range cvalues {
		total += val
	}
	return total
}

func GetCalibrationValue(line string) int {
	var first rune
	var last rune
	var isFirstAssigned = false
  var buff []rune
  var numberWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, rv := range line {
    if unicode.IsDigit(rune(rv)) {
			if isFirstAssigned == false {
				first = rv
				isFirstAssigned = true
			}
			last = rv
    } else {
      buff = append(buff, rv)
      for _, word := range numberWords {
        if strings.Contains(string(buff), word) {
          v, _ := lookup(word)
          if isFirstAssigned == false {
            first = v
            isFirstAssigned = true
          }
          last = v
          // There's a case where eighthree should be parsed as 83.
          // Clear down all but the last char in the buffer to workaround this.
          // Its icky but works.
          buff = buff[len(string(buff))-1:]
        } 
      }
    }
	}

	num, err := strconv.Atoi(fmt.Sprintf("%c%c", first, last))
  if err != nil {
    fmt.Println(err)
  }
  return num
}
