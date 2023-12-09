package main

import (
 "testing"
)


func TestGetCalibValue(t *testing.T) {
  tests := []struct {
    name string 
    val string 
    want int
  } {
    { "test two1nine", "two1nine", 29 },
    { "eightwothree", "eightwothree" , 83 },
    { "eighthree", "eighthree" , 83 },
    { "abcone2threexyz", "abcone2threexyz" , 13 },
    { "xtwone3four", "xtwone3four" , 24 },
    { "4nineeightseven2", "4nineeightseven2" , 42 },
    { "zoneight234", "zoneight234" , 14 },
    { "7pqrstsixteen", "7pqrstsixteen" , 76 },
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      got := GetCalibrationValue(tt.val)
      if got != tt.want {
        t.Errorf("GetCalibrationValue(%s) = %d want %d", tt.val, got, tt.want)
      }
    })
  }
}
