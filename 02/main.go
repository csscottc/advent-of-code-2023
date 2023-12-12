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
  var sumPower int
  var games []*Game 

  for {
		line, err := reader.ReadString('\n')
    if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error reading line:", err)
			}
			break
		}
    games = append(games,ParseLine(line))
    
  }

  for _, game := range games {
    var possibleGame bool = false
    for _, round := range game.rounds {
      d := round.isPossible()
      if d == false {
        possibleGame = false
        break;
      }
      possibleGame = true
    } 
    if possibleGame {
      total += game.id
    }

    lpRound := game.lpRound()
    power := lpRound.red * lpRound.green * lpRound.blue
    sumPower += power
  }


  fmt.Printf("(Part 1) The total is %d\n", total)
  fmt.Printf("(Part 2) The total power is %d\n", sumPower)
  

}

type Game struct {
  id int
  rounds []*Round
}

func (g *Game) lpRound() *Round {
  var highestRed int = 0
  var highestBlue int = 0
  var highestGreen int = 0
  for _, v := range g.rounds {
    if (v.red > highestRed) {
      highestRed = v.red
    }
    if (v.blue > highestBlue) {
      highestBlue = v.blue
    }
    if (v.green > highestGreen) {
      highestGreen = v.green
    }
  }

  return &Round{
    red: highestRed,
    blue: highestBlue,
    green: highestGreen,
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

func ParseLine(line string) *Game {
  gametag := strings.Split(line, ":")[0]
  rounds := strings.Split(strings.Split(line, ":")[1], ";")
  var pRounds []*Round
  for _, val := range rounds {
    pRounds = append(pRounds, ParseRound(strings.Trim(val," ")))
  }
  gameId, _ := strconv.Atoi(strings.Split(gametag," ")[1])
  return &Game { id: gameId, rounds: pRounds }
}
