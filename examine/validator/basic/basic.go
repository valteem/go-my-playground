package basic

type Person struct {
	GivenName  string `validate:"required"`
	FamilyName string `validate:"required"`
	Age        uint   `validate:"gte=0,lte=123"`
	Email      string `validate:"required,email"`
}
