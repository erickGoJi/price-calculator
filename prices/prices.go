package prices

import (
	"bufio"
	"fmt"
	"os"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (j TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file!")
		fmt.Println(err)
		file.Close()
		return
	}

}

// Process
func (j TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)
	// Calculate the price of a product with tax
	for _, price := range j.InputPrices {
		// Calculate the price of a product with tax
		result[fmt.Sprintf("%.2f", price)] = price * (1 + j.TaxRate)
	}

	fmt.Println(result)
}

// NewTaxIncludedPriceJob creates a new TaxIncludedPriceJob
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}

}
