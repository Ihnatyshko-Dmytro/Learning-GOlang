package main

import ( "strings" )

func High(s string) (hiWord string) {
    words := strings.Split(s, " ")
    hiPoints := 0
  
    for _, word := range words {
      totalPoint := 0
      
      for _, leter := range word {
        totalPoint += int(leter - 'a') + 1
      }
      if hiPoints < totalPoint {
        hiPoints = totalPoint
        hiWord = word
      }
    }
    return
}