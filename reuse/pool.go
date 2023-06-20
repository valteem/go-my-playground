package reuse

type Resource struct {
	Name string 
}

func New() any {
	return &Resource{"resource"}
}