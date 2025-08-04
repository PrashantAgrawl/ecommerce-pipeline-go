package pipeline

import (
	"encoding/csv"
	"os"
	"sort"
	"strconv"
)

func ReadRawSalesCSV(path string) ([]SaleRecord, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, _ := reader.ReadAll()

	var sales []SaleRecord
	for i, row := range rows {
		if i == 0 {
			continue
		}
		qty, _ := strconv.Atoi(row[3])
		price, _ := strconv.ParseFloat(row[4], 64)
		discount, _ := strconv.ParseFloat(row[5], 64)
		revenue, _ := strconv.ParseFloat(row[9], 64)

		sales = append(sales, SaleRecord{
			OrderID:         row[0],
			ProductName:     row[1],
			Category:        row[2],
			Quantity:        qty,
			UnitPrice:       price,
			DiscountPercent: discount,
			Region:          row[6],
			SaleDate:        row[7],
			CustomerEmail:   row[8],
			Revenue:         revenue,
		})
	}
	return sales, nil
}

func TopProducts(sales []SaleRecord) []ProductStats {
	stats := map[string]*ProductStats{}
	for _, sale := range sales {
		if _, ok := stats[sale.ProductName]; !ok {
			stats[sale.ProductName] = &ProductStats{Name: sale.ProductName}
		}
		stats[sale.ProductName].Revenue += sale.Revenue
		stats[sale.ProductName].Units += sale.Quantity
	}

	var sorted []ProductStats
	for _, v := range stats {
		sorted = append(sorted, *v)
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Revenue > sorted[j].Revenue
	})

	if len(sorted) > 10 {
		return sorted[:10]
	}
	return sorted
}

func WriteTopProductsToCSV(products []ProductStats, outputPath string) {
	f, _ := os.Create(outputPath)
	w := csv.NewWriter(f)
	w.Write([]string{"product_name", "revenue", "units_sold"})
	for _, p := range products {
		w.Write([]string{p.Name, strconv.FormatFloat(p.Revenue, 'f', 2, 64), strconv.Itoa(p.Units)})
	}
	w.Flush()
	f.Close()
}

func GenerateAllMetrics(rawPath string, outTop, outMonthly, outRegion, outAnomaly, outCat, outRPO, outRepeat string) error {
	sales, err := ReadRawSalesCSV(rawPath)
	if err != nil {
		return err
	}

	var valid []SaleRecord
	for _, s := range sales {
		if s.Quantity > 0 && s.UnitPrice > 0 && s.DiscountPercent >= 0 && s.DiscountPercent <= 1 {
			valid = append(valid, s)
		}
	}

	top := TopProducts(valid)
	WriteTopProductsToCSV(top, outTop)

	// Other metrics can be similarly written (see earlier logic)

	return nil
}
