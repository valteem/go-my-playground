package accesscontrol

import (
	"testing"
)

func TestEvaluate(t *testing.T) {

	testPermissions := map[string][]string{
		"create": {"datasource:userspace:*"},
		"view":   {"datasource:*", "external:userspace:*"},
	}

	tests := []struct {
		description string
		input       PermissionEvaluator
		output      bool
	}{
		{
			description: "action not permitted",
			input: PermissionEvaluator{
				Action: "delete",
			},
			output: false,
		},
		{
			description: "action permitted, scope not permitted",
			input: PermissionEvaluator{
				Action: "create",
				Scopes: []string{"external:userspace:*"},
			},
			output: false,
		},
		{
			description: `action "create" and scope are permitted`,
			input: PermissionEvaluator{
				Action: "create",
				Scopes: []string{"datasource:userspace:*"},
			},
			output: true,
		},
		{
			description: `action "view" and scope are permitted`,
			input: PermissionEvaluator{
				Action: "view",
				Scopes: []string{"external:userspace:*"},
			},
			output: true,
		},
	}

	for _, tc := range tests {
		if output := tc.input.Evaluate(testPermissions); output != tc.output {
			t.Errorf("%q: get %t, expect %t", tc.description, output, tc.output)
		}
	}

}
