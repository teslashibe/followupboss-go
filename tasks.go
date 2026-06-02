package followupboss

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// TasksQuery filters GET /tasks. Zero-valued fields are omitted.
type TasksQuery struct {
	AssignedUserID int
	PersonID       int
	IsCompleted    *bool // nil = no filter
	Sort           string
	Limit          int
	Offset         int
}

func (q TasksQuery) values() url.Values {
	v := url.Values{}
	setStr(v, "sort", q.Sort)
	if q.AssignedUserID > 0 {
		v.Set("assignedUserId", strconv.Itoa(q.AssignedUserID))
	}
	if q.PersonID > 0 {
		v.Set("personId", strconv.Itoa(q.PersonID))
	}
	if q.IsCompleted != nil {
		v.Set("isCompleted", strconv.FormatBool(*q.IsCompleted))
	}
	if q.Limit > 0 {
		v.Set("limit", strconv.Itoa(q.Limit))
	}
	if q.Offset > 0 {
		v.Set("offset", strconv.Itoa(q.Offset))
	}
	return v
}

// ListTasks returns tasks.
func (c *Client) ListTasks(ctx context.Context, q TasksQuery) (map[string]any, error) {
	return c.get(ctx, "/tasks", q.values())
}

// GetTask fetches a single task by id.
func (c *Client) GetTask(ctx context.Context, id int) (map[string]any, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: task id required", ErrInvalidParams)
	}
	return c.get(ctx, "/tasks/"+strconv.Itoa(id), nil)
}

// CompleteTask marks a task complete via PUT /tasks/{id} {"isCompleted": true}.
func (c *Client) CompleteTask(ctx context.Context, id int) (map[string]any, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: task id required", ErrInvalidParams)
	}
	return c.send(ctx, http.MethodPut, "/tasks/"+strconv.Itoa(id), map[string]any{"isCompleted": true})
}
