package followupboss

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// DealsQuery filters GET /deals. Zero-valued fields are omitted.
type DealsQuery struct {
	Stage      string
	PipelineID int
	Sort       string
	Limit      int
	Offset     int
}

func (q DealsQuery) values() url.Values {
	v := url.Values{}
	setStr(v, "stage", q.Stage)
	setStr(v, "sort", q.Sort)
	if q.PipelineID > 0 {
		v.Set("pipelineId", strconv.Itoa(q.PipelineID))
	}
	if q.Limit > 0 {
		v.Set("limit", strconv.Itoa(q.Limit))
	}
	if q.Offset > 0 {
		v.Set("offset", strconv.Itoa(q.Offset))
	}
	return v
}

// ListDeals returns deals (pipeline opportunities).
func (c *Client) ListDeals(ctx context.Context, q DealsQuery) (map[string]any, error) {
	return c.get(ctx, "/deals", q.values())
}

// GetDeal fetches a single deal by id.
func (c *Client) GetDeal(ctx context.Context, id int) (map[string]any, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: deal id required", ErrInvalidParams)
	}
	return c.get(ctx, "/deals/"+strconv.Itoa(id), nil)
}
