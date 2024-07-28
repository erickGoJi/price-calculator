package prices

import (
	"bufio"
	"fmt"
	"os"

	"example.com/price-calculator/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (j *TaxIncludedPriceJob) LoadData() {
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

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println("error converting strings to floats!")
		fmt.Println(err)
		file.Close()
		return
	}

	j.InputPrices = prices
	file.Close()
}

// Process
func (j *TaxIncludedPriceJob) Process() {
	j.LoadData()
	result := make(map[string]string)
	// Calculate the price of a product with tax
	for _, price := range j.InputPrices {
		taxIncludedPrice := price * (1 + j.TaxRate)
		// Calculate the price of a product with tax
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
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
