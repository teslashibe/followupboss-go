// Package followupboss is a small, dependency-free Go client + MCP tool
// surface for the Follow Up Boss REST API (https://api.followupboss.com/v1).
//
// Authentication
//
// Follow Up Boss uses HTTP Basic auth: the user's API key is the username and
// the password is left blank (https://docs.followupboss.com/reference/authentication).
// Obtain a key from Admin → API in the Follow Up Boss web app. An agent key
// only sees that agent's assigned people; an owner/broker key sees the whole
// account.
package followupboss

import (
	"net/http"
	"strings"
	"time"
)

const (
	baseURL          = "https://api.followupboss.com/v1"
	defaultUserAgent = "followupboss-go/0.1 (+https://github.com/teslashibe/followupboss-go)"
	defaultTimeout   = 30 * time.Second
	defaultRetries   = 3
	defaultRetryBase = 500 * time.Millisecond
)

// Client talks to the Follow Up Boss API.
type Client struct {
	apiKey     string
	httpClient *http.Client
	userAgent  string

	// Optional system identification headers. Follow Up Boss asks
	// integrations that act on behalf of customers to register and send
	// X-System / X-System-Key. Personal API-key use does not require them.
	system    string
	systemKey string

	maxRetries int
	retryBase  time.Duration
}

// Option configures a Client.
type Option func(*Client)

// WithHTTPClient overrides the default *http.Client. Nil is ignored.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		if hc != nil {
			c.httpClient = hc
		}
	}
}

// WithUserAgent overrides the default User-Agent.
func WithUserAgent(ua string) Option {
	return func(c *Client) {
		if ua != "" {
			c.userAgent = ua
		}
	}
}

// WithRetry sets the retry policy for 429 / 5xx responses.
func WithRetry(maxRetries int, base time.Duration) Option {
	return func(c *Client) {
		c.maxRetries = maxRetries
		c.retryBase = base
	}
}

// WithSystem sets the X-System / X-System-Key identification headers used by
// registered integrations.
func WithSystem(system, systemKey string) Option {
	return func(c *Client) {
		c.system = system
		c.systemKey = systemKey
	}
}

// New constructs a Client. The API key is required.
func New(apiKey string, opts ...Option) (*Client, error) {
	if strings.TrimSpace(apiKey) == "" {
		return nil, ErrInvalidAuth
	}
	c := &Client{
		apiKey:     strings.TrimSpace(apiKey),
		httpClient: &http.Client{Timeout: defaultTimeout},
		userAgent:  defaultUserAgent,
		maxRetries: defaultRetries,
		retryBase:  defaultRetryBase,
	}
	for _, o := range opts {
		o(c)
	}
	return c, nil
}
