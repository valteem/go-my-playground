// Need separate package for errors to avoid circular imports

package repoerr

import (
	"errors"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)
