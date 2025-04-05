package accesscontrol

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
			// TODO: add partial scope and wildcard matching
			if target == scope {
				return true
			}
		}
	}

	return false

}
