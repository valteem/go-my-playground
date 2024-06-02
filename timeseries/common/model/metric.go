package model

var (
	NameValidationScheme = LegacyValidation
)

type ValidationScheme int

const (
	LegacyValidation ValidationScheme = iota
	UTF8Validation
)

type EscapingScheme int

const (
	NoEscaping EscapingScheme = iota
	UnderscoreEscaping
	DotsEscaping
	ValueEncodingEscaping
)
