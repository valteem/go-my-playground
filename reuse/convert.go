package reuse

import "fmt"

func ConvertParamToSliceOfString(param []interface{}) []string {
	s := make([]string, len(param))
	for i, p := range param {
		// if v, ok := p.(string); !ok {
		// 	s[i] = "invalid"
		// } else {
		// 	s[i] = v
		// }
		s[i] = fmt.Sprintln(p)
	}
	return s
}