package mcp

import (
	"context"

	followupboss "github.com/teslashibe/followupboss-go"
	"github.com/teslashibe/mcptool"
)

// ListDealsInput is the typed input for followupboss_list_deals.
type ListDealsInput struct {
	Stage      string `json:"stage,omitempty" jsonschema:"description=Filter by deal stage name"`
	PipelineID int    `json:"pipeline_id,omitempty" jsonschema:"description=Filter by pipeline ID"`
	Sort       string `json:"sort,omitempty" jsonschema:"description=Sort order"`
	Limit      int    `json:"limit,omitempty" jsonschema:"description=Max results,minimum=1,maximum=100"`
	Offset     int    `json:"offset,omitempty" jsonschema:"description=Rows to skip,minimum=0"`
}

func listDeals(ctx context.Context, c *followupboss.Client, in ListDealsInput) (any, error) {
	return c.ListDeals(ctx, followupboss.DealsQuery{
		Stage:      in.Stage,
		PipelineID: in.PipelineID,
		Sort:       in.Sort,
		Limit:      in.Limit,
		Offset:     in.Offset,
	})
}

// GetDealInput is the typed input for followupboss_get_deal.
type GetDealInput struct {
	ID int `json:"id" jsonschema:"description=Deal ID,required"`
}

func getDeal(ctx context.Context, c *followupboss.Client, in GetDealInput) (any, error) {
	return c.GetDeal(ctx, in.ID)
}

var dealTools = []mcptool.Tool{
	mcptool.Define[*followupboss.Client, ListDealsInput](
		"followupboss_list_deals",
		"List Follow Up Boss deals (pipeline opportunities), optionally filtered by stage or pipeline.",
		"ListDeals",
		listDeals,
	),
	mcptool.Define[*followupboss.Client, GetDealInput](
		"followupboss_get_deal",
		"Fetch a single Follow Up Boss deal by ID.",
		"GetDeal",
		getDeal,
	),
}
