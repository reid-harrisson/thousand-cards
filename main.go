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

// Function to check if any subset of divisors sums to the card number
func subsetSum(divisors []int, target int) bool {
	dp := make([]bool, target+1)
	dp[0] = true

	for _, divisor := range divisors {
		for j := target; j >= divisor; j-- {
			dp[j] = dp[j] || dp[j-divisor]
		}
	}
	return dp[target]
}

func main() {
	validCards := []int{}

	for card := 1; card <= 1000; card++ {
		divisors := calculateDivisors(card)
		sumOfDivisors := 0
		for _, divisor := range divisors {
			sumOfDivisors += divisor
		}

		if sumOfDivisors > card && !subsetSum(divisors, card) {
			validCards = append(validCards, card)
		}
	}

	fmt.Println("Valid card numbers:", validCards)
}
