package mcp

import (
	"context"

	followupboss "github.com/teslashibe/followupboss-go"
	"github.com/teslashibe/mcptool"
)

// ListTasksInput is the typed input for followupboss_list_tasks.
type ListTasksInput struct {
	AssignedUserID int    `json:"assigned_user_id,omitempty" jsonschema:"description=Filter by assigned user ID"`
	PersonID       int    `json:"person_id,omitempty" jsonschema:"description=Filter by associated person ID"`
	IncludeDone    bool   `json:"include_done,omitempty" jsonschema:"description=Include completed tasks (default false)"`
	Sort           string `json:"sort,omitempty" jsonschema:"description=Sort order"`
	Limit          int    `json:"limit,omitempty" jsonschema:"description=Max results,minimum=1,maximum=100"`
	Offset         int    `json:"offset,omitempty" jsonschema:"description=Rows to skip,minimum=0"`
}

func listTasks(ctx context.Context, c *followupboss.Client, in ListTasksInput) (any, error) {
	q := followupboss.TasksQuery{
		AssignedUserID: in.AssignedUserID,
		PersonID:       in.PersonID,
		Sort:           in.Sort,
		Limit:          in.Limit,
		Offset:         in.Offset,
	}
	if !in.IncludeDone {
		done := false
		q.IsCompleted = &done
	}
	return c.ListTasks(ctx, q)
}

// CompleteTaskInput is the typed input for followupboss_complete_task.
type CompleteTaskInput struct {
	ID int `json:"id" jsonschema:"description=Task ID to mark complete,required"`
}

func completeTask(ctx context.Context, c *followupboss.Client, in CompleteTaskInput) (any, error) {
	return c.CompleteTask(ctx, in.ID)
}

var taskTools = []mcptool.Tool{
	mcptool.Define[*followupboss.Client, ListTasksInput](
		"followupboss_list_tasks",
		"List Follow Up Boss tasks, optionally filtered by user or person; excludes done by default.",
		"ListTasks",
		listTasks,
	),
	mcptool.Define[*followupboss.Client, CompleteTaskInput](
		"followupboss_complete_task",
		"Mark a Follow Up Boss task complete by ID.",
		"CompleteTask",
		completeTask,
	),
}
