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
- **Integration Events** - Integration Events API (create, acknowledge, close alerts, add notes)
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
go get github.com/circleyu/go-jsmops/v2@v2.1.0
```

## Authentication

This library supports two authentication methods:

1. **Basic Authentication**: For regular APIs (`/jsm/ops/api/{cloudId}/v1/...`)
   - Uses `userName` (email) and `apiToken` (API Token)
   - Applies to all standard JSM Operations APIs

2. **API Integration (GenieKey)**: For Integration Events API (`/jsm/ops/integration/v2/...`)
   - Uses `apiKey` (GenieKey)
   - Only for Integration Events API (create, acknowledge, close alerts, add notes)
   - Optional: Pass empty string if you don't need Integration Events API

## Quick Start

### Basic Usage (Basic Auth Only)

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/circleyu/go-jsmops/v2"
    "github.com/circleyu/go-jsmops/v2/alert"
    "github.com/sirupsen/logrus"
)

func main() {
    // Initialize client (Basic Auth only)
    logger := logrus.New()
    logger.SetLevel(logrus.DebugLevel)
    
    options := &jsmops.ClientOptions{
        Level:  jsmops.LogDebug,
        Logger: logger,
    }
    
    client := jsmops.Init(
        "your-cloud-id",
        "your-api-token",  // API Token for Basic Auth
        "your-email@example.com",  // Username/Email for Basic Auth
        "",  // apiKey (empty string means Integration Events API is not used)
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

### Using Integration Events API

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/circleyu/go-jsmops/v2"
    "github.com/circleyu/go-jsmops/v2/alert"
    "github.com/sirupsen/logrus"
)

func main() {
    // Initialize client (supports both Basic Auth and Integration Events API)
    logger := logrus.New()
    logger.SetLevel(logrus.DebugLevel)
    
    options := &jsmops.ClientOptions{
        Level:  jsmops.LogDebug,
        Logger: logger,
    }
    
    client := jsmops.Init(
        "your-cloud-id",
        "your-api-token",  // API Token for Basic Auth
        "your-email@example.com",  // Username/Email for Basic Auth
        "your-genie-key",  // GenieKey for Integration Events API
        options,
    )
    
    // Use regular API (uses Basic Auth)
    listReq := &alert.ListAlertsRequest{
        Limit: 10,
    }
    
    result, err := client.Alert.ListAlerts(listReq)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Found %d alerts\n", len(result.Alerts))
    
    // Use Integration Events API (uses GenieKey)
    createReq := &alert.IntegrationCreateAlertRequest{
        Message: "CPU usage exceeded 80%",
        Details: map[string]string{
            "server": "server-01",
            "cpu":    "85%",
        },
        Priority: alert.P1,
        Source:   "MonitoringTool",
    }
    
    createResult, err := client.IntegrationEvents.CreateAlert(createReq)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Alert creation request submitted: %s\n", createResult.RequestID)
    
    // Example: Acknowledge alert
    ackReq := &alert.IntegrationAcknowledgeAlertRequest{
        IdentifierType:  alert.ALERTID,
        IdentifierValue: "alert-id-here",
        User:            "John Smith",
        Source:          "MonitoringTool",
        Note:            "Working on this alert",
    }
    
    ackResult, err := client.IntegrationEvents.AcknowledgeAlert(ackReq)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Alert acknowledge request submitted: %s\n", ackResult.RequestID)
    
    // Example: Close alert
    closeReq := &alert.IntegrationCloseAlertRequest{
        IdentifierType:  alert.ALERTID,
        IdentifierValue: "alert-id-here",
        User:            "John Smith",
        Source:          "MonitoringTool",
        Note:            "Issue resolved",
    }
    
    closeResult, err := client.IntegrationEvents.CloseAlert(closeReq)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Alert close request submitted: %s\n", closeResult.RequestID)
    
    // Example: Add note
    noteReq := &alert.IntegrationAddNoteRequest{
        IdentifierType:  alert.ALERTID,
        IdentifierValue: "alert-id-here",
        User:            "John Smith",
        Source:          "MonitoringTool",
        Note:            "This is a note added via Integration Events API",
    }
    
    noteResult, err := client.IntegrationEvents.AddNote(noteReq)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Note add request submitted: %s\n", noteResult.RequestID)
}
```

### Without Logging

```go
client := jsmops.Init(
    "your-cloud-id",
    "your-api-token",
    "your-email@example.com",
    "",  // apiKey (empty string means Integration Events API is not used)
    jsmops.EmptyOptions(),
)
```

## Integration Events API

Integration Events API is a special set of API endpoints designed for creating and managing alerts from external systems (such as monitoring tools). These APIs use GenieKey authentication and do not require cloudID.

### Supported Operations

- **CreateAlert** - Create a new alert
- **AcknowledgeAlert** - Acknowledge an alert
- **CloseAlert** - Close an alert
- **AddNote** - Add a note to an alert

### Key Differences

The main differences between Integration Events API and regular Alerts API:

1. **Authentication**: Uses GenieKey instead of Basic Auth
2. **Path**: `/jsm/ops/integration/v2/...` instead of `/jsm/ops/api/{cloudId}/v1/...`
3. **Request Structure**:
   - `IntegrationCreateAlertRequest` uses `details` (`map[string]string`) instead of `extraProperties` (`map[string]interface{}`)
   - All requests include fields like `user`, `source`, `note`
4. **Response**: Returns 202 Accepted, indicating the request has been submitted for processing (asynchronous operation)

### Use Cases

Integration Events API is particularly suitable for:
- Monitoring systems automatically creating alerts
- CI/CD tools reporting build failures
- Automated scripts managing alert lifecycle
- Third-party integration systems

## Examples

Detailed usage examples can be found in the [examples](./examples) directory:

- [Basic Operations Example](./examples/basic/main.go) - Basic CRUD operations
- [Alert Management Example](./examples/alerts/main.go) - Alert creation, querying, and updates
- [Schedule Management Example](./examples/schedules/main.go) - Schedule and on-call management
- [Integration Management Example](./examples/integrations/main.go) - Integration and action management

## API Documentation

### Authentication

This library supports two authentication methods:

#### Basic Authentication (for regular APIs)

All standard JSM Operations APIs use Basic Authentication. You need to provide:
- `cloudID`: Jira Cloud ID
- `apiToken`: API Token (obtained from [Atlassian Account Settings](https://id.atlassian.com/manage-profile/security/api-tokens))
- `userName`: Your email address

To obtain an API Token:
1. Go to [Atlassian Account Settings](https://id.atlassian.com/manage-profile/security/api-tokens)
2. Click **Create API token**
3. Copy the generated API Token

#### API Integration (GenieKey) (for Integration Events API)

Integration Events API uses GenieKey authentication. You need to provide:
- `apiKey`: API Integration API Key (GenieKey)

To obtain a GenieKey, set up an API Integration in Jira Service Management:
1. Go to your team's Operations page
2. Select **Integrations** > **Add integration**
3. Search and select "API"
4. Enter an integration name and complete the setup
5. Copy the generated API Key

The API Key format is a UUID, for example: `g4ff854d-a14c-46a8-b8f0-0960774319dd`

**Note**: If you don't need to use Integration Events API, you can pass an empty string for the `apiKey` parameter.

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
├── main.go                        # Main client and initialization
├── http.go                        # HTTP request handling
├── endpoints.go                   # API endpoint definitions
├── alerts.go                      # Alert management
├── integration_events.go          # Integration Events API management
├── contacts.go                    # Contact management
├── ...                            # Other resource managers
├── alert/                         # Alert-related structures
│   ├── integration_events_request.go  # Integration Events API request structures
│   └── ...                        # Other alert-related structures
├── contacts/                      # Contact-related structures
└── examples/                      # Usage examples
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

## Version History

### v2.1.0
- ✅ Restored Basic Authentication for regular APIs
- ✅ Added Integration Events API support (using GenieKey authentication)
- ✅ Supports 4 Integration Events API endpoints: CreateAlert, AcknowledgeAlert, CloseAlert, AddNote

### v2.0.0
- ✅ Updated module path to `/v2` for Go module versioning support
- ✅ Switched to API Integration authentication (GenieKey)

## Related Links

- [Jira Service Management Operations API Documentation](https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-operations/)
- [Integration Events API Documentation](https://developer.atlassian.com/cloud/jira/service-desk-ops/rest/v1/api-group-integration-events/)

