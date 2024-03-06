package reuse

// https://www.digitalocean.com/community/tutorials/creating-custom-errors-in-go

type CustomError struct{}

func (e CustomError) Error() string {
	return "custom error"
}
