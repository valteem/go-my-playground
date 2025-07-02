package nested

type FoodCode uint32

const (
	Apples FoodCode = iota
	Cherries
	Onions

	NotFound = "not found"
)

var (
	foodStore = map[FoodCode]string{
		Apples:   "apples",
		Cherries: "cherries",
		Onions:   "onions",
	}
)

func FoodName(code FoodCode) string {
	if v, ok := foodStore[code]; ok {
		return v
	}
	return NotFound
}
