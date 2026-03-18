package main

import ("strconv")

func NextHigher(n int) int {
  
  result := 0
  startNum := n
  startBits := calculateBits(intToBinary(startNum))
  for result == 0 {
    startNum++
	secondNum := calculateBits(intToBinary(startNum))
    if  secondNum == startBits {
		result = startNum
	}
  }
  
  
  
  
  return result
}

func calculateBits (binaryNum string) int {
  numBits := 0
  
  for _, num := range binaryNum {
    if num == '1' {
      numBits++
    }
  }
  return numBits
}

func intToBinary (number int) string {
  binaryN := strconv.FormatInt(int64(number), 2)
  return binaryN
}