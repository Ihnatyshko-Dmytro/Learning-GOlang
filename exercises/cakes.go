package main

import "math"

func Cakes(recipe, available map[string]int) int {
	minCakes := math.MaxInt32 

	for ingredient, neededAmount := range recipe {
		availableAmount, ok := available[ingredient]
		
		if !ok || availableAmount < neededAmount {
			return 0
		}

		possibleWithThis := availableAmount / neededAmount

		if possibleWithThis < minCakes {
			minCakes = possibleWithThis
		}
	}

	return minCakes
}