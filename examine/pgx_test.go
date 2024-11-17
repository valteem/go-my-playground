package examine

import (
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestIdentifierSanitize(t *testing.T) {

	tests := []struct {
		input  pgx.Identifier
		output string
	}{
		{pgx.Identifier{"schema", "table"}, "\"schema\".\"table\""},
		{pgx.Identifier{"fruit", "apple"}, "\"fruit\".\"apple\""},
		{pgx.Identifier{"life", "is", "beach"}, "\"life\".\"is\".\"beach\""},
	}

	for _, tc := range tests {
		output := tc.input.Sanitize()
		if output != tc.output {
			t.Errorf("sanitizing %v:\nget\n%v\nexpect\n%s\n", tc.input, output, tc.output)
		}
	}
}
