package main

import "math"

func Cakes(recipe, available map[string]int) int {
	minCakes := math.MaxInt32 // Починаємо з дуже великого числа

	for ingredient, neededAmount := range recipe {
		availableAmount, ok := available[ingredient]
		
		// Якщо інгредієнта немає в наявності
		if !ok || availableAmount < neededAmount {
			return 0
		}

		// Скільки тортів можна спекти з цього конкретного інгредієнта
		possibleWithThis := availableAmount / neededAmount

		// Оновлюємо мінімум
		if possibleWithThis < minCakes {
			minCakes = possibleWithThis
		}
	}

	return minCakes
}