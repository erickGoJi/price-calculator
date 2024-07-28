package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (j *TaxIncludedPriceJob) LoadData() {
	lines, err := j.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println("error converting strings to floats!")
		fmt.Println(err)
		return
	}

	j.InputPrices = prices
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

	j.TaxIncludedPrices = result

	j.IOManager.WriteJSON(j)
}

// NewTaxIncludedPriceJob creates a new TaxIncludedPriceJob
func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}

}
