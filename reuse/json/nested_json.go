package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

var (
	storage ItemStorage
)

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type OrderRow struct {
	Item     *Item `json:"item"`
	Quantity int   `json:"quantity"`
}

type rawOrderRow struct {
	Item     string `json:"item"`
	Quantity string `json:"quantity"`
}

type Order struct {
	Order string      `json:"order"`
	Rows  []*OrderRow `json:"rows"`
}

func (o *Order) UnmarshalJSON(b []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	id, ok := m["order"]
	if !ok {
		return errors.New("order ID not found")
	}
	var orderID string
	err = json.Unmarshal(*id, &orderID)
	if err != nil {
		return err
	}
	o.Order = orderID

	rows, ok := m["rows"]
	if !ok {
		return errors.New("order rows not found")
	}
	var r []*json.RawMessage
	err = json.Unmarshal(*rows, &r)
	if err != nil {
		return err
	}
	for _, row := range r {
		ror := &rawOrderRow{}
		err = json.Unmarshal(*row, ror)
		if err != nil {
			return err
		}
		item, ok := storage.Get(ror.Item)
		if !ok {
			return fmt.Errorf("item %q not found", ror.Item)
		}
		qty, err := strconv.Atoi(ror.Quantity) // converting string to int - actually not required for proper JSON
		if err != nil {
			return fmt.Errorf("incorrect quantity format: %q", ror.Quantity)
		}
		o.Rows = append(o.Rows, &OrderRow{Item: item, Quantity: qty})
	}

	return nil
}

type OrderNoPreFill struct {
	Order string      `json:"order"`
	Rows  []*OrderRow `json:"rows"`
}
