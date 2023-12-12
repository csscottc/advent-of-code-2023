package main

import (
  "testing"
)

func roundsEqual(a, b *Round) bool {
    return a.green == b.green && a.red == b.red && a.blue == b.blue
}

func slicesRoundsEqual(slice1, slice2 []*Round) bool {
    if len(slice1) != len(slice2) {
        return false // Slices of different lengths are not equal
    }
    for i := range slice1 {
        if !roundsEqual(slice1[i], slice2[i]) {
            return false
        }
    }
    return true
}

func TestParseLine(t *testing.T) {
  tests := []struct {
    line string
    wantGameId int
    wantRounds []*Round 
  } {
    {
      line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
      wantRounds: []*Round{{blue:3, red: 4, green: 0}, {red: 1, green: 2, blue:6 }, {green: 2, blue: 0, red: 0},},
      // []string{"3 blue, 4 red", "1 red, 2 green, 6 blue", "2 green"},
      wantGameId: 1,
    },
  }

  for _, tt := range tests {
    t.Run(tt.line, func(t *testing.T) {
      game := ParseLine(tt.line)
      if game.id != tt.wantGameId {
        t.Errorf("Gametag mismatch got: %d want: %d", game.id, tt.wantGameId)
      }
      if !slicesRoundsEqual(game.rounds, tt.wantRounds) {
        t.Errorf("rounds didn't match wantRounds want: '%+v' got: '%+v'", tt.wantRounds, game.rounds)
      }
    })
  }

}

func TestParseRound(t *testing.T) {
  tests := []struct {
    name string
    round string
    want Round
  } {
    { name: "test 1", round: "3 blue, 4 red;", want: Round{ red: 4, blue: 3, green: 0 } },
    { name: "test 2", round: "1 blue, 4 red, 1 green;", want: Round{ red: 4, blue: 1, green: 1 } },
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) { 
      got := ParseRound(tt.round)
      if *got != tt.want {
        t.Errorf("TestParseRound in %s - want %+v out %+v", tt.round, tt.want, got)   
      }
    })
  }
}
