package model

import "regexp"

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

const (
	// Defines how metric and label names are to be escaped. Set in Accept or Content-Type header.
	EscapingKey = "escaping"

	// Possible values for EscapingKey:
	AllowUTF8         = "allow-utf-8" // no escaping
	EscapeUnderscores = "underscores"
	EscapeDots        = "dots"
	EscapeValues      = "values"
)

// Valid metric names
var MetricNameRE = regexp.MustCompile("`^[a-zA-Z_:][a-zA-Z0-9_:]*$`")
