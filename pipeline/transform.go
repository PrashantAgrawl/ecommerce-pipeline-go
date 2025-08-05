package pipeline

func GenerateAllMetrics(rawPath string,
	outTop string,
	outMonthly string,
	outRegion string,
	outAnomaly string,
	outCat string,
	outRPO string,
	outRepeat string) error {

	sales := []SaleRecord{}

	// Ingest and clean
	err := ReadCSVInChunks(rawPath, 100000, func(chunk [][]string) {
		for _, row := range chunk {
			if record, ok := CleanRecord(row); ok {
				sales = append(sales, record)
			}
		}
	})
	if err != nil {
		return err
	}

	// Metric Writers
	WriteTopProducts(sales, outTop)
	WriteMonthlySummary(sales, outMonthly)
	WriteRegionWisePerformance(sales, outRegion)
	WriteAnomalyRecords(sales, outAnomaly)
	WriteCategoryDiscountMap(sales, outCat)
	WriteRevenuePerOrder(sales, outRPO)
	WriteRepeatCustomerStats(sales, outRepeat)

	return nil
}
