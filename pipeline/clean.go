package pipeline

import (
	"strconv"
	"strings"
)

func CleanRecord(row []string) (SaleRecord, bool) {
	qty, err1 := strconv.Atoi(row[3])
	price, err2 := strconv.ParseFloat(row[4], 64)
	discount, err3 := strconv.ParseFloat(row[5], 64)
	revenue, err4 := strconv.ParseFloat(row[9], 64)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return SaleRecord{}, false
	}

	if qty <= 0 || discount < 0 || discount > 1 {
		return SaleRecord{}, false
	}

	region := strings.ToLower(row[6])
	regionMap := map[string]string{
		"nort": "North", "north": "North", "south": "South",
		"east": "East", "west": "West", "central": "Central",
	}
	if val, ok := regionMap[region]; ok {
		region = val
	}

	return SaleRecord{
		OrderID:         row[0],
		ProductName:     strings.TrimSpace(strings.ToLower(row[1])),
		Category:        strings.TrimSpace(strings.ToLower(row[2])),
		Quantity:        qty,
		UnitPrice:       price,
		DiscountPercent: discount,
		Region:          region,
		SaleDate:        row[7],
		CustomerEmail:   row[8],
		Revenue:         revenue,
	}, true
}
