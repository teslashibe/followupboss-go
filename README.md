# followupboss-go

A small, dependency-free Go client and [MCP](https://github.com/teslashibe/mcptool)
tool surface for the [Follow Up Boss](https://www.followupboss.com/) REST API
(`https://api.followupboss.com/v1`).

## Authentication

Follow Up Boss uses **HTTP Basic auth**: your API key is the username and the
password is blank. Get a key from **Admin → API** in the Follow Up Boss web app.

- An **agent** key only sees that agent's assigned people.
- An **owner/broker** key sees the whole account.

See the [auth docs](https://docs.followupboss.com/reference/authentication).

## Install

```bash
go get github.com/teslashibe/followupboss-go
```

## Client usage

```go
c, err := followupboss.New(os.Getenv("FUB_API_KEY"))
if err != nil {
    log.Fatal(err)
}

ctx := context.Background()

// Search contacts
people, _ := c.ListPeople(ctx, followupboss.PeopleQuery{Name: "drew", Limit: 25})

// Fetch one contact with all fields
person, _ := c.GetPerson(ctx, 10763, "allFields")

// Register a lead / log an activity
c.CreateEvent(ctx, followupboss.Event{
    Source:  "My Website",
    Type:    "Registration",
    Message: "Interested in 123 Main St",
    Person: followupboss.EventPerson{
        FirstName: "Tom",
        LastName:  "Minch",
        Emails:    followupboss.Emails("tom@example.com"),
        Phones:    followupboss.Phones("555-555-1234"),
    },
})

// Deals and tasks
deals, _ := c.ListDeals(ctx, followupboss.DealsQuery{Limit: 25})
tasks, _ := c.ListTasks(ctx, followupboss.TasksQuery{Limit: 25})
c.CompleteTask(ctx, 4242)
```

All read methods return the raw API envelope as `map[string]any` (e.g.
`{_metadata, people:[...]}`), keeping the full response available to callers.

## MCP tools

The `mcp` subpackage implements `mcptool.Provider` (platform `followupboss`):

| Tool | Wraps |
| --- | --- |
| `followupboss_list_people` | `ListPeople` |
| `followupboss_get_person` | `GetPerson` |
| `followupboss_list_deals` | `ListDeals` |
| `followupboss_get_deal` | `GetDeal` |
| `followupboss_list_tasks` | `ListTasks` |
| `followupboss_complete_task` | `CompleteTask` |
| `followupboss_create_event` | `CreateEvent` |

```go
provider := mcp.Provider{}
tools := provider.Tools()
```

## Demo

```bash
FUB_API_KEY=xxxx go run ./cmd/fub-demo
```

## License

MIT
