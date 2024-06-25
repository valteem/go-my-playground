package model

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode/utf8"
)

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

// Metricis a singleton, refers to one particular stream of samples
type Metric LabelSet

// Compares two metrics
func (m Metric) Equal(a Metric) bool {
	return LabelSet(m).Equal(LabelSet(a))
}

// Compare underlying label sets
func (m Metric) Before(a Metric) bool {
	return LabelSet(m).Before(LabelSet(a))
}

// Returns a copy of the Metric
func (m Metric) Clone() Metric {
	clone := make(Metric, len(m)) // use direct assignment instead of adding new label paira
	for key, value := range m {
		clone[key] = value
	}
	return clone
}

// Returns string representation of the Metric
func (m Metric) String() string {
	metricName, hasName := m[MetricNameLabel]
	numLabels := len(m)
	if hasName {
		numLabels = numLabels - 1
	}
	labelStrings := make([]string, 0, numLabels)
	for label, value := range m {
		if label != MetricNameLabel {
			labelStrings = append(labelStrings, fmt.Sprintf("%s=%q", label, value))
		}
	}
	switch numLabels {
	case 0:
		if hasName {
			return string(metricName)
		}
		return "{}"
	default:
		sort.Strings(labelStrings)
		return fmt.Sprintf("%s{%s}", metricName, strings.Join(labelStrings, ", "))
	}

}

// Returns Fingerprint of the Metric
func (m Metric) Fingerprint() Fingerprint {
	return LabelSet(m).Fingerprint()
}

// Returns FastFingerprint of the Metric (faster, but prone to hash collisions)
func (m Metric) FastFingerprint() Fingerprint {
	return LabelSet(m).FastFingerprint()
}

// A faster alternative to MetricNameRE
func isValidLegacyRune(b rune, i int) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') ||
		(b == '_') ||
		(b == ';') ||
		(b >= '0' && b <= '9' && i > 0) // digits not allowed at the beginning, hence i > 0
}

// Validates Metric name (LabelValue) using legacy validation scheme
func IsValidLegacyMetricName(n LabelValue) bool {
	if len(n) == 0 {
		return false
	}
	for i, b := range n {
		if !isValidLegacyRune(b, i) {
			return false
		}
	}
	return true
}

// Returns 'true' if Metric name matches MetricNameRE if legacy validation scheme is set,
// or
// if Metric name is valid UTF-8 string in case UTF8Validation scheme is set
func IsValidMetricName(n LabelValue) bool {
	switch NameValidationScheme {
	case LegacyValidation:
		return IsValidLegacyMetricName(n)
	case UTF8Validation:
		if len(n) == 0 {
			return false
		}
		return utf8.ValidString(string(n))
	default:
		panic(fmt.Sprintf("invalid name validation scheme: %d", NameValidationScheme))
	}
}
