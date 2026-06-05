package mcp

import (
	"context"

	followupboss "github.com/teslashibe/followupboss-go"
	"github.com/teslashibe/mcptool"
)

// ListPeopleInput is the typed input for followupboss_list_people.
type ListPeopleInput struct {
	Name      string `json:"name,omitempty" jsonschema:"description=Partial name match (e.g. drew matches Andrew)"`
	Email     string `json:"email,omitempty" jsonschema:"description=Exact email to search for"`
	Phone     string `json:"phone,omitempty" jsonschema:"description=Phone number to search for"`
	Stage     string `json:"stage,omitempty" jsonschema:"description=Stage name (e.g. Lead Trash)"`
	Tags      string `json:"tags,omitempty" jsonschema:"description=Comma-separated tags; matches any"`
	Sort      string `json:"sort,omitempty" jsonschema:"description=Sort order (default created)"`
	Fields    string `json:"fields,omitempty" jsonschema:"description=Comma-separated projection fields (e.g. emails phones firstName lastName stage tags allFields). NOTE email/phone are SEARCH-only top-level filters use the email/phone params for those not here"`
	Limit     int    `json:"limit,omitempty" jsonschema:"description=Max results 1-100,minimum=1,maximum=100"`
	Offset    int    `json:"offset,omitempty" jsonschema:"description=Rows to skip,minimum=0"`
}

func listPeople(ctx context.Context, c *followupboss.Client, in ListPeopleInput) (any, error) {
	return c.ListPeople(ctx, followupboss.PeopleQuery{
		Name:   in.Name,
		Email:  in.Email,
		Phone:  in.Phone,
		Stage:  in.Stage,
		Tags:   in.Tags,
		Sort:   in.Sort,
		Fields: in.Fields,
		Limit:  in.Limit,
		Offset: in.Offset,
	})
}

// GetPersonInput is the typed input for followupboss_get_person.
type GetPersonInput struct {
	ID     int    `json:"id" jsonschema:"description=Person ID,required"`
	Fields string `json:"fields,omitempty" jsonschema:"description=Comma-separated projection fields (e.g. emails phones firstName lastName stage tags allFields). NOTE email/phone are SEARCH-only filters not projection names"`
}

func getPerson(ctx context.Context, c *followupboss.Client, in GetPersonInput) (any, error) {
	return c.GetPerson(ctx, in.ID, in.Fields)
}

var peopleTools = []mcptool.Tool{
	mcptool.Define[*followupboss.Client, ListPeopleInput](
		"followupboss_list_people",
		"Search Follow Up Boss contacts by name, email, phone, stage, or tags.",
		"ListPeople",
		listPeople,
	),
	mcptool.Define[*followupboss.Client, GetPersonInput](
		"followupboss_get_person",
		"Fetch a single Follow Up Boss contact by ID, with optional field selection.",
		"GetPerson",
		getPerson,
	),
}
