package followupboss

import "context"

// HealthCheck verifies the API key is valid by hitting the lightweight
// identity endpoint (GET /me). Returns nil when the credential is live, or a
// sentinel error (e.g. ErrUnauthorized / ErrForbidden) when it is not.
//
// Hosts can type-assert any client to
// `interface{ HealthCheck(context.Context) error }` to probe credential
// liveness without importing this package.
func (c *Client) HealthCheck(ctx context.Context) error {
	_, err := c.Identity(ctx)
	return err
}
