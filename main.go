package main

import "fmt"

// Function Price calculates the price of a product with tax
func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)

	// Calculate the price of a product with tax
	for _, taxRate := range taxRates {
		// Create a slice to store the prices with tax
		var taxIncludedPrices []float64 = make([]float64, len(prices))
		// Calculate the price of a product with tax
		for priceIndex, price := range prices {
			// Calculate the price of a product with tax
			taxIncludedPrices[priceIndex] = price * (1 + taxRate)
		}
		// Store the prices with tax in the result map
		result[taxRate] = taxIncludedPrices
	}

	fmt.Println(result)
}
