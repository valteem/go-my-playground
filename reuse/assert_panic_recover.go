package reuse

import "fmt"

func CheckTypePerson(i interface{}) (p *Person, err error) {
	defer func() {
		if v:= recover(); v != nil {
			p = &Person{}
			err = fmt.Errorf("%s", v)
		}
	}()
	a := i.(*Person)
	return a, nil
}