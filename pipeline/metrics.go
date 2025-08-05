package pipeline

import (
	"encoding/csv"
	"os"
	"sort"
	"strconv"
)

func WriteTopProducts(sales []SaleRecord, outputPath string) {
	stats := map[string]*ProductStats{}
	for _, s := range sales {
		if _, ok := stats[s.ProductName]; !ok {
			stats[s.ProductName] = &ProductStats{Name: s.ProductName}
		}
		stats[s.ProductName].Revenue += s.Revenue
		stats[s.ProductName].Units += s.Quantity
	}

	var sorted []ProductStats
	for _, v := range stats {
		sorted = append(sorted, *v)
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Revenue > sorted[j].Revenue
	})

	if len(sorted) > 10 {
		sorted = sorted[:10]
	}

	f, _ := os.Create(outputPath)
	w := csv.NewWriter(f)
	w.Write([]string{"product_name", "revenue", "units_sold"})
	for _, p := range sorted {
		w.Write([]string{p.Name, strconv.FormatFloat(p.Revenue, 'f', 2, 64), strconv.Itoa(p.Units)})
	}
	w.Flush()
	f.Close()
}

func WriteMonthlySummary(sales []SaleRecord, outputPath string) {
	monthly := map[string][3]float64{}
	count := map[string]int{}
	for _, s := range sales {
		month := ""
		if len(s.SaleDate) >= 7 {
			month = s.SaleDate[:7]
		}
		agg := monthly[month]
		agg[0] += s.Revenue
		agg[1] += float64(s.Quantity)
		agg[2] += s.DiscountPercent
		count[month]++
		monthly[month] = agg
	}

	f, _ := os.Create(outputPath)
	w := csv.NewWriter(f)
	w.Write([]string{"month", "revenue", "quantity", "avg_discount"})
	for k, v := range monthly {
		c := float64(count[k])
		w.Write([]string{k, strconv.FormatFloat(v[0], 'f', 2, 64), strconv.Itoa(int(v[1])), strconv.FormatFloat(v[2]/c, 'f', 2, 64)})
	}
	w.Flush()
	f.Close()
}

func WriteRegionWisePerformance(sales []SaleRecord, outputPath string) {
	regions := map[string][3]float64{}
	count := map[string]int{}
	for _, s := range sales {
		agg := regions[s.Region]
		agg[0] += s.Revenue
		agg[1] += float64(s.Quantity)
		agg[2] += s.DiscountPercent
		count[s.Region]++
		regions[s.Region] = agg
	}

	f, _ := os.Create(outputPath)
	w := csv.NewWriter(f)
	w.Write([]string{"region", "revenue", "quantity", "avg_discount"})
	for k, v := range regions {
		c := float64(count[k])
		w.Write([]string{k, strconv.FormatFloat(v[0], 'f', 2, 64), strconv.Itoa(int(v[1])), strconv.FormatFloat(v[2]/c, 'f', 2, 64)})
	}
	w.Flush()
	f.Close()
}

func WriteAnomalyRecords(sales []SaleRecord, outputPath string) {
	sort.Slice(sales, func(i, j int) bool {
		return sales[i].Revenue > sales[j].Revenue
	})

	f, _ := os.Create(outputPath)
	w := csv.NewWriter(f)
	w.Write([]string{"order_id", "product_name", "category", "quantity", "unit_price", "discount_percent", "region", "sale_date", "customer_email", "revenue"})
	for i := 0; i < 5 && i < len(sales); i++ {
		s := sales[i]
		w.Write([]string{
			s.OrderID, s.ProductName, s.Category,
			strconv.Itoa(s.Quantity),
			strconv.FormatFloat(s.UnitPrice, 'f', 2, 64),
			strconv.FormatFloat(s.DiscountPercent, 'f', 2, 64),
			s.Region, s.SaleDate, s.CustomerEmail,
			strconv.FormatFloat(s.Revenue, 'f', 2, 64),
		})
	}
	w.Flush()
	f.Close()
}

func WriteCategoryDiscountMap(sales []SaleRecord, outputPath string) {
	catMap := map[string][2]float64{}
	count := map[string]int{}
	for _, s := range sales {
		agg := catMap[s.Category]
		agg[0] += s.DiscountPercent
		agg[1] += s.Revenue
		count[s.Category]++
		catMap[s.Category] = agg
	}

	f, _ := os.Create(outputPath)
	w := csv.NewWriter(f)
	w.Write([]string{"category", "avg_discount", "revenue"})
	for k, v := range catMap {
		c := float64(count[k])
		w.Write([]string{
			k,
			strconv.FormatFloat(v[0]/c, 'f', 2, 64),
			strconv.FormatFloat(v[1], 'f', 2, 64),
		})
	}
	w.Flush()
	f.Close()
}

func WriteRevenuePerOrder(sales []SaleRecord, outputPath string) {
	totalRevenue := 0.0
	orderSet := map[string]bool{}
	for _, s := range sales {
		totalRevenue += s.Revenue
		orderSet[s.OrderID] = true
	}
	orderCount := len(orderSet)
	avgRevenue := totalRevenue / float64(orderCount)

	f, _ := os.Create(outputPath)
	w := csv.NewWriter(f)
	w.Write([]string{"metric", "value"})
	w.Write([]string{"total_revenue", strconv.FormatFloat(totalRevenue, 'f', 2, 64)})
	w.Write([]string{"total_orders", strconv.Itoa(orderCount)})
	w.Write([]string{"avg_revenue_per_order", strconv.FormatFloat(avgRevenue, 'f', 2, 64)})
	w.Flush()
	f.Close()
}

func WriteRepeatCustomerStats(sales []SaleRecord, outputPath string) {
	emailMap := map[string]int{}
	for _, s := range sales {
		if s.CustomerEmail != "" {
			emailMap[s.CustomerEmail]++
		}
	}

	typeCount := map[string][2]float64{}
	for _, s := range sales {
		if s.CustomerEmail == "" {
			continue
		}
		cType := "New"
		if emailMap[s.CustomerEmail] > 1 {
			cType = "Repeat"
		}
		agg := typeCount[cType]
		agg[0] += 1
		agg[1] += s.Revenue
		typeCount[cType] = agg
	}

	f, _ := os.Create(outputPath)
	w := csv.NewWriter(f)
	w.Write([]string{"customer_type", "order_count", "revenue"})
	for k, v := range typeCount {
		w.Write([]string{k, strconv.Itoa(int(v[0])), strconv.FormatFloat(v[1], 'f', 2, 64)})
	}
	w.Flush()
	f.Close()
}
