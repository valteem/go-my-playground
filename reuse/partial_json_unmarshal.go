package reuse

import (
	"encoding/json"
	"fmt"
)

type AircraftFeatures struct {
	Model string `json:"model"`
	Registration string `json:"registration"`
	Age int `json:"age"` 
}

func ReadAircraftFeatures(b *[]byte) (a *AircraftFeatures, err error) {
	
	var af AircraftFeatures

	var objmap map[string]*json.RawMessage

	err = json.Unmarshal(*b, &objmap)
    if err != nil {
		return &AircraftFeatures{}, err
    }

	err = json.Unmarshal(*objmap["aircraft"], &af)
	if err != nil {
		return &AircraftFeatures{}, err
	}

	return &af, nil

}

func UnmarshalRawJSON(b *[]byte) (err error) {
	
	var objmap map[string]*json.RawMessage

	err = json.Unmarshal(*b, &objmap)

    if err != nil {
    	fmt.Println(err)
		return err
    }

	var flightid, flight_time string
	err = json.Unmarshal(*objmap["flightid"], &flightid) // pointer to 'map value' (*json.RawMessage)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(*objmap["flight_time"], &flight_time)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}