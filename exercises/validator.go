package main

import (
  "strings"
  "strconv"
)

func Is_valid_ip(ip string) bool {
  numbers := strings.Split(ip, ".")
  
  if len(numbers) != 4 {
    return false
  }
  for _, octet := range numbers {
    if len(octet) > 1 && octet[0] == '0' {
        return false
    }
    val, err := strconv.Atoi(octet)
    if err != nil || val < 0 || val > 255 {
      return false
    }
    if strconv.Itoa(val) != octet  {
      return false
    }
  }
  return true
}
