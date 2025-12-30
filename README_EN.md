# go-jsmops

A simple and easy-to-use Go library for calling the Jira Service Management Operations REST API.

[中文](README.md) | [English](README_EN.md)

## Features

- ✅ Complete API Coverage: Supports all Jira Service Management Operations API endpoints
- ✅ Type Safety: Uses strongly-typed Go structs for requests and responses
- ✅ Easy to Use: Clear interface design following Go conventions
- ✅ Logging Support: Optional logging functionality for debugging
- ✅ Error Handling: Unified error handling mechanism

## Supported Resource Categories

- **Alerts** - Alert management
- **Audit Logs** - Audit logs
- **Contacts** - Contact management
- **Teams** - Team management
- **Roles** - Custom user roles
- **Escalations** - Escalation rules
- **Forwarding Rules** - Forwarding rules
- **Heartbeats** - Heartbeat monitoring
- **Integrations** - Integration management
- **Integration Actions** - Integration actions
- **Integration Filters** - Integration filters
- **Maintenances** - Maintenance plans (global and team)
- **Notification Rules** - Notification rules
- **Notification Rule Steps** - Notification rule steps
- **Policies** - Alert policies (global and team)
- **Team Roles** - Team roles
- **Routing Rules** - Routing rules
- **Schedules** - Schedule management
- **Schedule On-calls** - Schedule on-calls
- **Schedule Overrides** - Schedule overrides
- **Schedule Rotations** - Schedule rotations
- **Schedule Timelines** - Schedule timelines
- **Syncs** - Sync management
- **Sync Actions** - Sync actions
- **Sync Action Groups** - Sync action groups
- **JEC** - JEC channel management

## Installation

```bash
go get github.com/circleyu/go-jsmops
```

## Quick Start

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/circleyu/go-jsmops"
    "github.com/circleyu/go-jsmops/alert"
    "github.com/sirupsen/logrus"
)

func main() {
    // Initialize client
    logger := logrus.New()
    logger.SetLevel(logrus.DebugLevel)
    
    options := &jsmops.ClientOptions{
        Level:  jsmops.LogDebug,
        Logger: logger,
    }
    
    client := jsmops.Init(
        "your-cloud-id",
        "your-api-key",
        options,
    )
    
    // Use API
    // Example: List alerts
    listReq := &alert.ListAlertsRequest{
        Limit: 10,
    }
    
    result, err := client.Alert.ListAlerts(listReq)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Found %d alerts\n", len(result.Alerts))
}
```

### Without Logging

```go
client := jsmops.Init(
    "your-cloud-id",
    "your-api-key",
    jsmops.EmptyOptions(),
)
```

## Examples

Detailed usage examples can be found in the [examples](./examples) directory:

- [Basic Operations Example](./examples/basic/main.go) - Basic CRUD operations
- [Alert Management Example](./examples/alerts/main.go) - Alert creation, querying, and updates
- [Schedule Management Example](./examples/schedules/main.go) - Schedule and on-call management
- [Integration Management Example](./examples/integrations/main.go) - Integration and action management

## API Documentation

### Authentication

All API requests use API Integration authentication (GenieKey). You need to provide:
- `cloudID`: Jira Cloud ID
- `apiKey`: API Integration API Key

To obtain an API Key, set up an API Integration in Jira Service Management:
1. Go to your team's Operations page
2. Select **Integrations** > **Add integration**
3. Search and select "API"
4. Enter an integration name and complete the setup
5. Copy the generated API Key

The API Key format is a UUID, for example: `g4ff854d-a14c-46a8-b8f0-0960774319dd`

### Error Handling

All API methods return errors. The error type is `APIError`, which includes HTTP status code and error details.

```go
result, err := client.Alert.GetAlert(&alert.GetAlertRequest{ID: "alert-id"})
if err != nil {
    if apiErr, ok := err.(jsmops.APIError); ok {
        fmt.Printf("API Error: %s (Status Code: %d)\n", apiErr.Error(), apiErr.StatusCode)
    } else {
        fmt.Printf("Other Error: %v\n", err)
    }
}
```

## Development

### Project Structure

```
go-jsmops/
├── main.go              # Main client and initialization
├── http.go              # HTTP request handling
├── endpoints.go         # API endpoint definitions
├── alerts.go            # Alert management
├── contacts.go          # Contact management
├── ...                  # Other resource managers
├── alert/               # Alert-related structures
├── contacts/            # Contact-related structures
└── examples/            # Usage examples
```

### Adding New Features

1. Create `request.go` and `result.go` in the corresponding resource directory
2. Create or update the corresponding manager file (e.g., `alerts.go`)
3. Add endpoint definitions in `endpoints.go`
4. Register the new manager in `main.go`

## License

This project is licensed under the MIT License.

## Contributing

Issues and Pull Requests are welcome!

## Related Links

- [Jira Service Management Operations API Documentation](https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-operations/)

