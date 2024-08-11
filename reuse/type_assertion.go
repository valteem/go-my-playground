package reuse

type StockItem struct {
	ID          int
	Description string
	SupplierID  int
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

type TestObjectA struct {
	ID int
}

func (o TestObjectA) Show() int {
	return o.ID
}

type TestObjectB struct {
	ID int
}

func (o TestObjectB) Output() int {
	return o.ID
}

func TellObjectType(o any) string {
	switch o.(type) {
	case interface{ Show() int }:
		return "TypeObjectA"
	case interface{ Output() int }:
		return "TypeObjectB"
	default:
		return "Unknown"
	}
}
