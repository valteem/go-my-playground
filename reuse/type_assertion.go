package reuse

type StockItem struct {
	ID int
	Description string
	SupplierID int
}

func AssertType(v any) string {

	_, ok := v.(Person)
	if ok {
		return "Person" 
	}

	_, ok = v.(StockItem)
	if ok {
		return "Stock Item"
	}

	return "Unknown"
}

