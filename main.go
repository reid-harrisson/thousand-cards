package main

import "fmt"

// Function to calculate divisors of a number
func calculateDivisors(n int) ([]int, int) {
	divisors := []int{1}
	sumOfDivisors := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			sumOfDivisors += i
			if i != n/i { // Avoid adding the square root twice
				divisors = append(divisors, n/i)
				sumOfDivisors += n / i
			}
		}
	}
	return divisors, sumOfDivisors
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
		divisors, sumOfDivisors := calculateDivisors(card)

		if sumOfDivisors > card && !subsetSum(divisors, card) {
			validCards = append(validCards, card)
		}
	}

	fmt.Println("Valid card numbers:", validCards)
}
