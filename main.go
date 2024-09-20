package main

import "fmt"

func main() {
	transactions := []float64{}
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	balance := calculateBalance(transactions)
	fmt.Printf("Total balance: %.2f\n", balance)

}

func scanTransaction() float64 {
	var amount float64
	fmt.Print("Enter transaction amount: ")
	fmt.Scan(&amount)
	return amount
}

func calculateBalance(transactions []float64) float64 {
	balance := 0.0
	for _, transaction := range transactions {
		balance += transaction
	}
	return balance
}
