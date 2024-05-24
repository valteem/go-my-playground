package model

var (
	NameValidationScheme = LegacyValidation
)

type ValidationScheme int

const (
	LegacyValidation ValidationScheme = iota
	UTF8Validation
)
