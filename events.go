package followupboss

import (
	"context"
	"fmt"
	"net/http"
)

// Event is the payload for POST /events — the canonical way to register a
// lead or log an activity. Follow Up Boss dedupes/merges the embedded person
// by email/phone and routes via the account's lead-flow rules.
// See https://docs.followupboss.com/reference/events-post.
type Event struct {
	// Source is the lead source label (required), e.g. "My Website".
	Source string `json:"source"`
	// Type is the event type, e.g. "Registration", "Inquiry",
	// "Property Inquiry", "General Inquiry".
	Type string `json:"type,omitempty"`
	// Message is the free-text note associated with the event.
	Message string `json:"message,omitempty"`
	// Person is the contact this event is about.
	Person EventPerson `json:"person"`
}

// EventPerson is the contact embedded in an Event. Emails/phones are arrays
// of {value} objects in the API; helpers below build them.
type EventPerson struct {
	FirstName string              `json:"firstName,omitempty"`
	LastName  string              `json:"lastName,omitempty"`
	Emails    []map[string]string `json:"emails,omitempty"`
	Phones    []map[string]string `json:"phones,omitempty"`
	Tags      []string            `json:"tags,omitempty"`
}

// CreateEvent posts an event/lead. Source and at least one of the person's
// name/email/phone should be set.
func (c *Client) CreateEvent(ctx context.Context, e Event) (map[string]any, error) {
	if e.Source == "" {
		return nil, fmt.Errorf("%w: event source required", ErrInvalidParams)
	}
	return c.send(ctx, http.MethodPost, "/events", e)
}

// Emails builds the API's [{value}] shape from plain addresses.
func Emails(addrs ...string) []map[string]string {
	out := make([]map[string]string, 0, len(addrs))
	for _, a := range addrs {
		if a != "" {
			out = append(out, map[string]string{"value": a})
		}
	}
	return out
}

// Phones builds the API's [{value}] shape from plain numbers.
func Phones(nums ...string) []map[string]string {
	out := make([]map[string]string, 0, len(nums))
	for _, n := range nums {
		if n != "" {
			out = append(out, map[string]string{"value": n})
		}
	}
	return out
}
