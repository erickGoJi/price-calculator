package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
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
