package main	

import (
  "strings"
  "strconv"
)

func Points(games []string) int {
  totalPoints := 0
  
  for _, match := range games {
    parts := strings.Split(match, ":")
    
    x, _ := strconv.Atoi(parts[0])
    y, _ := strconv.Atoi(parts[1])
    
    switch {
      case x > y:
      totalPoints =+ 2
      case x == y:
      totalPoints =+ 1
    }
    
  }
  return totalPoints
}