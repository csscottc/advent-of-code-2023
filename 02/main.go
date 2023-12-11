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

  var total int
	
  for {
		line, err := reader.ReadString('\n')
    if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error reading line:", err)
			}
			break
		}
    game, rounds := ParseLine(line)
    var possibleGame bool = false
    for _, round := range rounds {
      d := ParseRound(round).isPossible()
      if d == false {
        possibleGame = false
        break;
      }
      possibleGame = true
    } 

    // var possibleGameIds []int = make([]int, 5)
    fmt.Printf("%d is possible: %t\n", game, possibleGame)
    if possibleGame {
      total += game
      //possibleGameIds = append(possibleGameIds, game)
    }
    fmt.Printf("%d\n",total)
    
  }
}

func (r *Round) isPossible() bool {
  if r.red > 12 {
    return false
  }
  if r.green > 13 {
    return false
  }
  if r.blue > 14 {
    return false
  }
  return true
}

type Round struct {
  red int
  blue int
  green int
}

func ParseRound(round string) *Round {
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

func ParseLine(line string) (int, []string) {
  gametag := strings.Split(line, ":")[0]
  rounds := strings.Split(strings.Split(line, ":")[1], ";")
  for i, val := range rounds {
    strings.Trim(val," ")
    rounds[i] = strings.Trim(val," ") 
  }
  gameId, _ := strconv.Atoi(strings.Split(gametag," ")[1])
  return gameId, rounds
}
