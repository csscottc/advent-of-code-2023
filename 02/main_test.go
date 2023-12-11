package main

import (
  "testing"
)

func compareSlices[T comparable](slice1, slice2 []T) bool {
    if len(slice1) != len(slice2) {
        return false
    }
    for i := range slice1 {
        if slice1[i] != slice2[i] {
            return false
        }
    }
    return true
}

func TestParseLine(t *testing.T) {
  tests := []struct {
    line string
    wantGameId string
    wantRounds []string 
  } {
    {
      line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
      wantRounds: []string{"3 blue, 4 red", "1 red, 2 green, 6 blue", "2 green"},
      wantGameId: 1,
    },
  }

  for _, tt := range tests {
    t.Run(tt.line, func(t *testing.T) {
      gotGameId, gotRounds := ParseLine(tt.line)
      if gotGameId != tt.wantGameId {
        t.Errorf("Gametag mismatch got: %s  want: %s", gotGameId, tt.wantGameId)
      }
      if !compareSlices(gotRounds, tt.wantRounds) {
        t.Errorf("gotRounds didn't match wantRounds want: '%+v' got: '%+v'", tt.wantRounds, gotRounds)
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
