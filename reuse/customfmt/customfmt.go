package customfmt

import (
	"fmt"
)

type Address struct {
	ZipCode string
	City    string
	Street  string
}

type Building struct {
	TotalArea      float32
	NumberOfFloors int
	Addr           *Address
}

func (b Building) GoString() string {
	if b.Addr != nil {
		addrStr := fmt.Sprintf("Address: zip code %s, city %s, street %s",
			b.Addr.ZipCode,
			b.Addr.City,
			b.Addr.Street)
		return fmt.Sprintf("Total area %.1f sq.m., number of floors %d, %s",
			b.TotalArea,
			b.NumberOfFloors,
			addrStr)
	}
	return fmt.Sprintf("Total area %f sq.m., number of floors %d",
		b.TotalArea,
		b.NumberOfFloors)
}
