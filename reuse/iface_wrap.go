// Wrap a type into interface and get it back with type assertion

package reuse

type Person struct {
	Name string
	Age int
}

func ConvertPersonToIface(p *Person) interface{} {
	var v interface{} = p
	return v
}

func ConvertIfaceToPerson(i interface{}) (p *Person) {
	a, ok := i.(*Person)
	if !ok {
		return &Person{}
	}
	return a
}