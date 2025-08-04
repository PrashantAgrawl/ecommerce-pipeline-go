package pipeline

type SaleRecord struct {
	OrderID         string
	ProductName     string
	Category        string
	Quantity        int
	UnitPrice       float64
	DiscountPercent float64
	Region          string
	SaleDate        string
	CustomerEmail   string
	Revenue         float64
}

type ProductStats struct {
	Name    string
	Revenue float64
	Units   int
}
