package accesscontrol

import "strings"

// Action is evaluated user action, scopes describe where action can be performed.
// For example:
// Action: rules:create
// Scopes: "folders:*", "folders:uid:*"
type PermissionEvaluator struct {
	Action string
	Scopes []string
}

// permissions are user permissions, user is obtained from request context
func (p PermissionEvaluator) Evaluate(permissions map[string][]string) bool {

	userScopes, ok := permissions[p.Action]
	if !ok {
		return false
	}

	// Not checking user scopes if action is found in permissions and no input scopes are presented (???)
	if len(p.Scopes) == 0 {
		return true
	}

	for _, target := range p.Scopes {
		for _, scope := range userScopes {
			if match(scope, target) {
				return true
			}
		}
	}

	return false

}

func ValidateScope(scope, target string) bool {
	prefix, last := scope[:len(scope)-1], scope[len(scope)-1]
	if len(prefix) > 0 && last == '*' {
		lastChar := prefix[len(prefix)-1] // last symbol before asterisk should be : or /
		if !(lastChar == ':' || lastChar == '/') {
			return false
		}
	}
	return !strings.ContainsAny(prefix, "?*")
}

func match(scope, target string) bool {

	if scope == "" {
		return false
	}

	if !ValidateScope(scope, target) {
		return false
	}

	prefix, last := scope[:len(scope)-1], scope[len(scope)-1]
	if last == '*' {
		if strings.HasPrefix(target, prefix) {
			return true
		}
	}

	return scope == target

}
