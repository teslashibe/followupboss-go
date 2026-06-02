// Command fub-demo is a tiny smoke test for the Follow Up Boss client.
//
// Usage:
//
//	FUB_API_KEY=xxxx go run ./cmd/fub-demo
//
// It prints the authenticated identity and the first page of people.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	followupboss "github.com/teslashibe/followupboss-go"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run() error {
	key := os.Getenv("FUB_API_KEY")
	if key == "" {
		return fmt.Errorf("set FUB_API_KEY")
	}
	c, err := followupboss.New(key)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	id, err := c.Identity(ctx)
	if err != nil {
		return fmt.Errorf("identity: %w", err)
	}
	dump("identity", id)

	people, err := c.ListPeople(ctx, followupboss.PeopleQuery{Limit: 5})
	if err != nil {
		return fmt.Errorf("list people: %w", err)
	}
	dump("people", people)
	return nil
}

func dump(label string, v any) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Printf("=== %s ===\n%s\n", label, b)
}
