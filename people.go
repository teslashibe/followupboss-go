package followupboss

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// normalizeFields maps Follow Up Boss search-filter tokens to their valid
// response-projection names within a comma-separated fields string. FUB treats
// "email"/"phone" as top-level SEARCH filters, not projection field names; the
// projection equivalents are plural ("emails"/"phones"). Other tokens and
// whitespace-trimmed casing are left untouched. Empty input returns "".
func normalizeFields(fields string) string {
	if fields == "" {
		return ""
	}
	parts := strings.Split(fields, ",")
	for i, p := range parts {
		switch strings.TrimSpace(p) {
		case "email":
			parts[i] = "emails"
		case "phone":
			parts[i] = "phones"
		default:
			parts[i] = strings.TrimSpace(p)
		}
	}
	return strings.Join(parts, ",")
}

// PeopleQuery filters GET /people. Zero-valued fields are omitted.
// See https://docs.followupboss.com/reference/people-get.
type PeopleQuery struct {
	Name         string // partial name match
	FirstName    string
	LastName     string
	Email        string
	Phone        string
	Stage        string // e.g. "Lead", "Trash"
	Source       string
	Tags         string // comma-separated; matches any
	AssignedTo   string
	Sort         string // default "created"
	Fields       string // comma-separated, or "allFields"/"allCustom"
	Limit        int    // max 100, default 10
	Offset       int
	IncludeTrash bool
}

func (q PeopleQuery) values() url.Values {
	v := url.Values{}
	setStr(v, "name", q.Name)
	setStr(v, "firstName", q.FirstName)
	setStr(v, "lastName", q.LastName)
	setStr(v, "email", q.Email)
	setStr(v, "phone", q.Phone)
	setStr(v, "stage", q.Stage)
	setStr(v, "source", q.Source)
	setStr(v, "tags", q.Tags)
	setStr(v, "assignedTo", q.AssignedTo)
	setStr(v, "sort", q.Sort)
	setStr(v, "fields", normalizeFields(q.Fields))
	if q.Limit > 0 {
		v.Set("limit", strconv.Itoa(q.Limit))
	}
	if q.Offset > 0 {
		v.Set("offset", strconv.Itoa(q.Offset))
	}
	if q.IncludeTrash {
		v.Set("includeTrash", "true")
	}
	return v
}

// ListPeople searches contacts. Returns the raw API envelope
// ({_metadata, people:[...]}).
func (c *Client) ListPeople(ctx context.Context, q PeopleQuery) (map[string]any, error) {
	return c.get(ctx, "/people", q.values())
}

// GetPerson fetches a single contact by id. fields is optional (comma-
// separated, or "allFields").
func (c *Client) GetPerson(ctx context.Context, id int, fields string) (map[string]any, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: person id required", ErrInvalidParams)
	}
	v := url.Values{}
	setStr(v, "fields", normalizeFields(fields))
	return c.get(ctx, "/people/"+strconv.Itoa(id), v)
}

// Identity returns the current API user/account (GET /me). Useful as a
// connection health check.
func (c *Client) Identity(ctx context.Context) (map[string]any, error) {
	return c.get(ctx, "/me", nil)
}

func setStr(v url.Values, key, val string) {
	if val != "" {
		v.Set(key, val)
	}
}
