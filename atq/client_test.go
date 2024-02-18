package atq

import (
	"testing"
)

func TestOptionTypes(t *testing.T) {
	value_check := (MaxRetryOpt == 1) &&
	               (QueueOpt == 2) &&
				   (TimeoutOpt == 3)
	if !value_check {
		t.Errorf("option types: wrong assignment")
	}
}