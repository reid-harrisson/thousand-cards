# Thousand Cards

## Overview

In this project, we explore a deck of one thousand cards numbered from 1 to 1000. The goal is to identify the cards that meet the following criteria:

1. The sum of its divisors (excluding the number itself) is greater than the card number.
2. No subset of those divisors sums up to the card number itself.

## Approach

To solve this problem, we will implement a Go program that performs the following steps:

1. **Calculate Divisors**: For each card number, calculate its divisors.
2. **Sum of Divisors**: Compute the sum of the divisors for each card.
3. **Check Conditions**: Verify if the sum of the divisors is greater than the card number and check if no subset of the divisors sums up to the card number.
4. **Output Valid Cards**: Collect and print the card numbers that satisfy both conditions.

## Implementation

The Go program is structured to efficiently calculate divisors and check the required conditions using a combination of loops and recursive functions.

### How to Execute

1. **Clone the Repository**:

   ```bash
   git clone <repository-url>
   cd thousand-cards
   ```

2. **Run the Program**:

   ```bash
   go run main.go
   ```

3. **Output**: The program will print the card numbers that meet the specified criteria.

## Conclusion

This project demonstrates the application of mathematical concepts in programming and provides a practical solution to the fortune teller's challenge. Feel free to explore and modify the code to enhance its functionality or performance.
