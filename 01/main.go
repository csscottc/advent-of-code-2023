package main

import (
	"bufio"
	"fmt"
	"os"
  "unicode"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {

	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error reading line:", err)
			}
			break
		}
		calibrationValue := getCalibrationValue(line)
		fmt.Println(calibrationValue);
	}
}

func getCalibrationValue(line string) string {
  var first rune;
  var last rune;
	for _,rv := range line {
    //fmt.Printf("%c %b \n", rv,unicode.IsDigit(rune(rv)))
    if unicode.IsDigit(rune(rv)) {
      if first == 0 {
        first = rv;
      }
      last = rv;
    }
	}
  // fmt.Printf("first:%c last:%c \n", first, last);
	return fmt.Sprintf("%c%c", first, last);
}
