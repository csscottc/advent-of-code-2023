package main

import (
  "fmt"
  "regexp"
  "strconv"
  "os"
  "bufio"
  "strings"
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
    ParseLine(line);
  }

}

type Round struct {
  red int
  blue int
  green int
}

func ParseRound(round string) *Round {
  fmt.Println("Parsing round", round)
  pattern := `(\d+)\s*([a-zA-Z]+)`
  re, err := regexp.Compile(pattern)
  if err != nil {
    fmt.Println("Error compiling regex:", err)
  }

  parsedRound := &Round{} 
  matches := re.FindAllStringSubmatch(round, -1)
  for _, m := range matches {
   v, err := strconv.Atoi(m[1])
   if (err == nil) {
     if(m[2] == "red") {
      parsedRound.red = v
     }
     if(m[2] == "blue") {
      parsedRound.blue = v
     }
     if(m[2] == "green") {
      parsedRound.green = v
     }
   } else {
    fmt.Printf("Error parsing value %s to integer", m[1])
   }
  }
  return parsedRound
}

func ParseLine(line string) (string, []string) {
  fmt.Println("Parsing line", line)
  gametag := strings.Split(line, ":")[0]
  rounds := strings.Split(strings.Split(line, ":")[1], ";")
  return gametag, rounds
}
