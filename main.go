package jsmops

import (
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/bytedance/sonic"
	"github.com/sirupsen/logrus"
)

// APIClient ...
type APIClient struct {
	cloudID               string
	userName              string
	apiToken              string
	apiKey                string // Optional, for Integration Events API
	logLevel              LogLevel
	logger                *logrus.Logger
	Alert                 AlertsManager
	AuditLogs             AuditLogsManager
	Contacts              ContactsManager
	Teams                 TeamsManager
	Roles                 RolesManager
	Escalations           EscalationsManager
	ForwardingRules       ForwardingRulesManager
	Heartbeats            HeartbeatsManager
	Integrations          IntegrationsManager
	IntegrationActions    IntegrationActionsManager
	IntegrationFilters    IntegrationFiltersManager
	Maintenances          MaintenancesManager
	NotificationRules     NotificationRulesManager
	NotificationRuleSteps NotificationRuleStepsManager
	Policies              PoliciesManager
	TeamPolicies          TeamPoliciesManager
	TeamRoles             TeamRolesManager
	RoutingRules          RoutingRulesManager
	Schedules             SchedulesManager
	SchedulesOnCalls      SchedulesOnCallsManager
	SchedulesOverrides    SchedulesOverridesManager
	SchedulesRotations    SchedulesRotationsManager
	SchedulesTimelines    SchedulesTimelinesManager
	Syncs                 SyncsManager
	SyncsActions          SyncsActionsManager
	SyncsActionGroups     SyncsActionGroupsManager
	JEC                   JECManager
	IntegrationEvents     IntegrationEventsManager
}

type LogLevel int

const (
	LogError LogLevel = iota
	LogInfo
	LogDebug
)

// ErrManagerNotRegistered is returned when a manager method is called but the required authentication is not provided
var ErrManagerNotRegistered = errors.New("manager not registered: authentication credentials not provided")

// ClientOptions ...
type ClientOptions struct {
	Level  LogLevel
	Logger *logrus.Logger
}

// EmptyOptions ...
func EmptyOptions() *ClientOptions {
	return nil
}

func NewOptions() *ClientOptions {
	return &ClientOptions{
		Level: LogError,
	}
}

// Init initializes the package
// cloudID: Jira Service Management cloud ID
// apiToken: API token for Basic Authentication (used for regular APIs)
// userName: Username/email for Basic Authentication (used for regular APIs)
// apiKey: Optional API key for Integration Events API (GenieKey). Empty string means Integration Events API is not used.
// Returns an error if neither Basic Auth nor GenieKey is provided.
func Init(cloudID, apiToken, userName string, apiKey string, options *ClientOptions) (*APIClient, error) {
	// Validate that at least one authentication method is provided
	hasBasicAuth := userName != "" && apiToken != ""
	hasGenieKey := apiKey != ""
	if !hasBasicAuth && !hasGenieKey {
		return nil, errors.New("at least one authentication method must be provided: either Basic Auth (userName and apiToken) or API Integration (apiKey)")
	}

	client := &APIClient{
		cloudID:  cloudID,
		userName: userName,
		apiToken: apiToken,
		apiKey:   apiKey,
	}
	if options != nil {
		client.logger = options.Logger
		client.logLevel = options.Level
	}

	// Update logging to reflect actual authentication methods
	if client.logger != nil {
		if hasBasicAuth && hasGenieKey {
			client.logger.Infof("JSM ops Client initializing..., authorization = Basic Auth (user: %s) and Integration Events API (GenieKey)", userName)
		} else if hasBasicAuth {
			client.logger.Infof("JSM ops Client initializing..., authorization = Basic Auth (user: %s)", userName)
		} else if hasGenieKey {
			client.logger.Infof("JSM ops Client initializing..., authorization = Integration Events API (GenieKey)")
		}
	}

	// Conditionally register standard JSM Operations API managers (only if Basic Auth is provided)
	if hasBasicAuth {
		client.Alert = newAlertsManager(client)
		client.AuditLogs = newAuditLogsManager(client)
		client.Contacts = newContactsManager(client)
		client.Teams = newTeamsManager(client)
		client.Roles = newRolesManager(client)
		client.Escalations = newEscalationsManager(client)
		client.ForwardingRules = newForwardingRulesManager(client)
		client.Heartbeats = newHeartbeatsManager(client)
		client.Integrations = newIntegrationsManager(client)
		client.IntegrationActions = newIntegrationActionsManager(client)
		client.IntegrationFilters = newIntegrationFiltersManager(client)
		client.Maintenances = newMaintenancesManager(client)
		client.NotificationRules = newNotificationRulesManager(client)
		client.NotificationRuleSteps = newNotificationRuleStepsManager(client)
		client.Policies = newPoliciesManager(client)
		client.TeamPolicies = newTeamPoliciesManager(client)
		client.TeamRoles = newTeamRolesManager(client)
		client.RoutingRules = newRoutingRulesManager(client)
		client.Schedules = newSchedulesManager(client)
		client.SchedulesOnCalls = newSchedulesOnCallsManager(client)
		client.SchedulesOverrides = newSchedulesOverridesManager(client)
		client.SchedulesRotations = newSchedulesRotationsManager(client)
		client.SchedulesTimelines = newSchedulesTimelinesManager(client)
		client.Syncs = newSyncsManager(client)
		client.SyncsActions = newSyncsActionsManager(client)
		client.SyncsActionGroups = newSyncsActionGroupsManager(client)
		client.JEC = newJECManager(client)
	}

	// Register Integration Events manager only if apiKey is provided
	if hasGenieKey {
		client.IntegrationEvents = newIntegrationEventsManager(client)
	}

	return client, nil
}

// func (client *APIClient) logErr(err error) {
// 	if err != nil && client.logger != nil {
// 		client.logger.Error(err.Error())
// 	}
// }

func (client *APIClient) logReq(req *http.Request) {
	if client.logger != nil {
		client.logger.Debug("Headers")
		for k, v := range req.Header {
			client.logger.Debugf("%s: %s\n", k, v)
		}
		client.logger.Debugf("URL: %s", req.URL)
		if req.Body != nil {
			body, _ := io.ReadAll(req.Body)
			// 使用 sonic 格式化 JSON
			var jsonData interface{}
			if err := sonic.Unmarshal(body, &jsonData); err == nil {
				if formatted, err := sonic.MarshalIndent(jsonData, "", "\t"); err == nil {
					client.logger.Debug(string(formatted))
				} else {
					client.logger.Debug(string(body))
				}
			} else {
				client.logger.Debug(string(body))
			}
		}
	}
}

func (client *APIClient) logRes(res *http.Response) {
	if client.logger != nil {
		client.logger.Debugf("Status: %d", res.StatusCode)
		if res.StatusCode != 200 {
			body, err := io.ReadAll(res.Body)
			if err != nil {
				client.logger.Debug(err.Error())
			}
			// 使用 sonic 格式化 JSON
			var jsonData interface{}
			if err := sonic.Unmarshal(body, &jsonData); err == nil {
				if formatted, err := sonic.MarshalIndent(jsonData, "", "\t"); err == nil {
					client.logger.Debug(string(formatted))
				} else {
					client.logger.Debug(string(body))
				}
			} else {
				client.logger.Debug(string(body))
			}
		}
	}
}

// BackupJSON ...
func (client *APIClient) BackupJSON(fileName string, data interface{}) error {
	backupJSON, _ := sonic.Marshal(data)
	return os.WriteFile(fileName, backupJSON, 0644)
}

// checkBasicAuth checks if Basic Authentication credentials are available
func (c *APIClient) checkBasicAuth() error {
	if c.userName == "" || c.apiToken == "" {
		return ErrManagerNotRegistered
	}
	return nil
}

// checkGenieKey checks if GenieKey (API Integration) credentials are available
func (c *APIClient) checkGenieKey() error {
	if c.apiKey == "" {
		return ErrManagerNotRegistered
	}
	return nil
}
