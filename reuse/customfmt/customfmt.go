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

// https://github.com/golang/go/issues/51195

type Help struct {
	Cmd string
}

func (h Help) Format(s fmt.State, r rune) {
	switch r {
	case 'h':
		fmt.Fprintf(s, "Help yourself: %s", h.Cmd)
	case 'm':
		fmt.Fprintf(s, "Maybe: %s", h.Cmd)
	default:
		fmtDir := fmt.FormatString(s, r)
		fmt.Fprintf(s, fmtDir, h.Cmd)
	}
}
