package model

var (
	NameValidationScheme = LegacyValidation
)

// Defines how metric and label names are validated
type ValidationScheme int

const (
	// Requires that metric and label names conform to MetricNameRE
	LegacyValidation ValidationScheme = iota
	// Requires that metric and label names are valid UTF-8 strings
	UTF8Validation
)

type EscapingScheme int

const (
	NoEscaping EscapingScheme = iota
	// Replaces all legacy-invalid (?) characters with underscores
	UnderscoreEscaping
	// Same as UnderscoreEscaping, except that dots are replace with `_dot_`, and single underscores converted to double underscores
	DotsEscaping
	// Prepends the name with `U_`, replaces all invalid (?) characters with unicode value, surrounded by underscores
	// Replaces single underscores with double underscores
	ValueEncodingEscaping
)
