package main

import (
  "testing"
)

func TestParseLine(t *testing.T) {
  tests := []struct {
    line string
    wantGametag string
    wantRounds []string 
  } {
    {
      line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
      wantRounds: []string{"3 blue, 4 red", "1 red, 2 green 6 blue", "2 green"},
      wantGametag: "Game 1",
    },
  }

  for _, tt := range tests {
    t.Run(tt.line, func(t *testing.T) {
      gotGametag, gotRounds := ParseLine(tt.line)
      if gotGametag != tt.wantGametag {
        t.Errorf("Gametag mismatch got: %s  want: %s", gotGametag, tt.wantGametag)
      }
      if gotRounds[0] != tt.wantRounds[0] {
        // TODO: I should write a comparison func
        t.Errorf("gotRounds[0] want: '%s' got: '%s'", tt.wantRounds[0], gotRounds[0])
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
