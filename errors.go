package followupboss

import (
	"errors"
	"fmt"
)

// Sentinel errors.
var (
	ErrInvalidAuth   = errors.New("followupboss: missing or invalid API key")
	ErrUnauthorized  = errors.New("followupboss: unauthorized (bad API key)")
	ErrForbidden     = errors.New("followupboss: forbidden (account locked or insufficient permission)")
	ErrNotFound      = errors.New("followupboss: not found")
	ErrRateLimited   = errors.New("followupboss: rate limited")
	ErrInvalidParams = errors.New("followupboss: invalid parameters")
	ErrRequestFailed = errors.New("followupboss: request failed")
)

// HTTPError is returned for unexpected non-2xx responses.
type HTTPError struct {
	StatusCode int
	Body       string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("followupboss: HTTP %d: %s", e.StatusCode, e.Body)
}
