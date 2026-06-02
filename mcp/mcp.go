// Package mcp exposes the Follow Up Boss client as a set of MCP tools.
package mcp

import "github.com/teslashibe/mcptool"

// Provider implements mcptool.Provider for Follow Up Boss.
type Provider struct{}

// Platform returns the credential-store identifier for this package.
func (Provider) Platform() string { return "followupboss" }

// Tools returns every Follow Up Boss MCP tool.
func (Provider) Tools() []mcptool.Tool {
	out := make([]mcptool.Tool, 0, len(peopleTools)+len(dealTools)+len(taskTools)+len(eventTools))
	out = append(out, peopleTools...)
	out = append(out, dealTools...)
	out = append(out, taskTools...)
	out = append(out, eventTools...)
	return out
}
