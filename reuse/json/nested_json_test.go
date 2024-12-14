package json

import (
	"encoding/json"
	"testing"

	"github.com/tidwall/gjson"
)

func TestOrderID(t *testing.T) {

	m := NewMemItemSorage()
	m.init(10)
	storage = m

	orderInput := `{"order":"0001","rows":[{"item":"i001","quantity":"1000"},{"item":"i002","quantity":"2000"}]}`

	order := Order{}

	err := json.Unmarshal([]byte(orderInput), &order)
	if err != nil {
		t.Fatalf("failed to unmarshal input string: %v", err)
	}

	total := 0
	for _, r := range order.Rows {
		total += r.Quantity
	}
	if total != 3000 {
		t.Errorf("total quantity: get %d, expect %d", total, 3000)
	}

	output, err := json.Marshal(order)
	if err != nil {
		t.Fatalf("failed to marshal output order: %v", err)
	}

	orderOutput := string(output)

	item := gjson.Get(orderOutput, "rows.#(quantity=\"1000\").item.name")
	itemExpected := `i001`
	if item.Str != itemExpected {
		t.Errorf("order items:\nget\n%s\nexpect\n%s", item.Str, itemExpected)
	}

}
