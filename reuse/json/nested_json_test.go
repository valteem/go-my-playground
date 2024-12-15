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

// do not need custom unmarshal for creating all nested objects
// instead of looking for existing ones (no pre-filled storage)
func TestUnmarshalStd(t *testing.T) {

	orderInput := `{"order":"0001","rows":[{"item":{"name":"i001", "description":"item #001"},"quantity":1000},{"item":{"name":"i002", "description":"item #002"},"quantity":2000}]}`

	order := OrderNoPreFill{}

	err := json.Unmarshal([]byte(orderInput), &order)
	if err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	if order.Order != "0001" {
		t.Errorf("order ID: get %q, expect %q", order.Order, "0001")
	}

	if len(order.Rows) != 2 {
		t.Errorf("number of rows in order: get %d, expect %d", len(order.Rows), 2)
	}

	if name, qty := order.Rows[0].Item.Name, order.Rows[0].Quantity; name != "i001" || qty != 1000 {
		t.Errorf("order row #0 (name, quantity): get (%q, %d), expect (%q, %d)", name, qty, "i001", 1000)
	}
	if name, qty := order.Rows[1].Item.Name, order.Rows[1].Quantity; name != "i002" || qty != 2000 {
		t.Errorf("order row #0 (name, quantity): get (%q, %d), expect (%q, %d)", name, qty, "i002", 2000)
	}

}
