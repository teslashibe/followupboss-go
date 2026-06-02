package mcp

import (
	"context"

	followupboss "github.com/teslashibe/followupboss-go"
	"github.com/teslashibe/mcptool"
)

// CreateEventInput is the typed input for followupboss_create_event. Posting
// an event registers/merges a lead and logs an activity.
type CreateEventInput struct {
	Source    string `json:"source" jsonschema:"description=Lead source label (e.g. My Website),required"`
	Type      string `json:"type,omitempty" jsonschema:"description=Event type (e.g. Registration Inquiry Property Inquiry)"`
	Message   string `json:"message,omitempty" jsonschema:"description=Free-text note for the event"`
	FirstName string `json:"first_name,omitempty" jsonschema:"description=Contact first name"`
	LastName  string `json:"last_name,omitempty" jsonschema:"description=Contact last name"`
	Email     string `json:"email,omitempty" jsonschema:"description=Contact email address"`
	Phone     string `json:"phone,omitempty" jsonschema:"description=Contact phone number"`
}

func createEvent(ctx context.Context, c *followupboss.Client, in CreateEventInput) (any, error) {
	return c.CreateEvent(ctx, followupboss.Event{
		Source:  in.Source,
		Type:    in.Type,
		Message: in.Message,
		Person: followupboss.EventPerson{
			FirstName: in.FirstName,
			LastName:  in.LastName,
			Emails:    followupboss.Emails(in.Email),
			Phones:    followupboss.Phones(in.Phone),
		},
	})
}

var eventTools = []mcptool.Tool{
	mcptool.Define[*followupboss.Client, CreateEventInput](
		"followupboss_create_event",
		"Create a Follow Up Boss event to register/merge a lead and log an activity.",
		"CreateEvent",
		createEvent,
	),
}
