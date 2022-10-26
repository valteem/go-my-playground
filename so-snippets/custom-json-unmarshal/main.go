// https://stackoverflow.com/questions/72095328/how-to-accommodate-for-2-different-types-of-the-same-field-when-parsing-json

package main

import (
	"encoding/json"
	"fmt"
)

type ApiResponse struct {
	Result struct {
		Value struct {
			Error ErrorMessage `json:"error"`
			Logs []string      `json:"logs"`
		} `json:"value"`
	} `json:"result"`
}

type ErrorMessage struct {
	Text string
}

func (e *ErrorMessage) UnmarshalJSON(data []byte) error {
	if (len(data)==0) || string(data) == "null" {
		return nil
	}
	if data[0] == '"' && data[len(data)-1] == '"' { // string ?
		return json.Unmarshal(data, &e.Text)
	}
	if data[0] == '{' && data[len(data)-1] == '}' { // object ?
		type T struct {
			Text string `json:"msg"`
		}
		return json.Unmarshal(data, (*T)(e)) // need clarification
	}
	return fmt.Errorf("unsupported error message type")
}

func main() {

	data1 := []byte(`{
            "result": {
			    "value": {
				    "error": "Invalid data",
				    "logs": []
    			}
            }
        }
	`)
	var r1 ApiResponse
	if err := json.Unmarshal(data1, &r1); err != nil {
		panic(err)
	}
	fmt.Println(r1.Result.Value.Error.Text)

	data2 := []byte(`{
	    	"result": {
		    	"value": {
			    	"error": {
				    	"msg": "Invalid data"
				    },
				    "logs": []
    			}
		    }
	    }
        `)
var r2 ApiResponse
if err := json.Unmarshal(data2, &r2); err != nil {
	panic(err)
}
fmt.Println(r1.Result.Value.Error.Text)

}