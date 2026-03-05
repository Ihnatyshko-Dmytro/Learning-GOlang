package main

func main () {
	Tribonacci([3]float64{20, 11, 6}, 0)
}

func Tribonacci(signature [3]float64, n int) []float64 {
  result := make([]float64, 0, n)
  for i := 0; i < n && i < 3; i++ {
        result = append(result, signature[i])
    }

  for i := 3; i < n; i++ {
    result = append(result, (result[i-3] + result[i-2] + result[i-1]))
  }
  return result
  
}