package main

import "fmt"

// Function to calculate divisors of a number
func calculateDivisors(n int) []int {
	divisors := []int{}
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func main() {
	for card := 1; card <= 1000; card++ {
		divisors := calculateDivisors(card)
		sumOfDivisors := 0
		for _, divisor := range divisors {
			sumOfDivisors += divisor
		}

		if sumOfDivisors > card {
			fmt.Println(card)
		}
	}
}
