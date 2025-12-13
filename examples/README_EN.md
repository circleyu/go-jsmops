# Usage Examples

This directory contains usage examples for the go-jsmops library.

[中文](README.md) | [English](README_EN.md)

## Environment Setup

Before running the examples, please set the following environment variables:

```bash
export JIRA_CLOUD_ID="your-cloud-id"
export JIRA_API_TOKEN="your-api-token"
export JIRA_USERNAME="your-username@example.com"
```

### Getting Jira Cloud ID

1. Log in to your Jira instance
2. Visit `https://admin.atlassian.com/`
3. Select your organization
4. Find the Cloud ID in the settings

### Getting API Token

1. Visit `https://id.atlassian.com/manage-profile/security/api-tokens`
2. Click "Create API token"
3. Copy the generated token

## Example Descriptions

### basic/main.go

Basic operations example, demonstrating how to:
- Initialize the client
- List teams
- List and get contacts

Run with:
```bash
cd examples/basic
go run main.go
```

### alerts/main.go

Alert management example, demonstrating how to:
- List alerts
- Create new alerts
- Get alert details
- Add and list alert notes

Run with:
```bash
cd examples/alerts
go run main.go
```

**Note**: Creating alerts requires a valid team ID. Please replace `your-team-id` in the code.

### schedules/main.go

Schedule management example, demonstrating how to:
- List schedules
- Get schedule details
- List on-call responders
- List next on-call responders

Run with:
```bash
cd examples/schedules
go run main.go
```

### integrations/main.go

Integration management example, demonstrating how to:
- List integrations
- Get integration details
- List integration actions
- Get integration alert filters

Run with:
```bash
cd examples/integrations
go run main.go
```

## Custom Examples

You can refer to these examples to create your own programs. Main steps:

1. Import necessary packages
2. Set environment variables or provide authentication information directly
3. Initialize the client
4. Use the corresponding manager to call API methods

## Error Handling

All API methods return errors. It is recommended to always check for errors:

```go
result, err := client.Alert.ListAlerts(req)
if err != nil {
    log.Fatalf("Operation failed: %v", err)
}
```

## More Resources

- [Main README](../README.md)
- [Jira Service Management Operations API Documentation](https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-operations/)

