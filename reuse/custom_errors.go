package reuse

// https://www.digitalocean.com/community/tutorials/creating-custom-errors-in-go

type CustomError struct{}

func (e CustomError) Error() string {
	return "custom error"
}

type AnotherCustomError struct {
	Err error
}

func (e AnotherCustomError) Error() string {
	return e.Err.Error()
}
