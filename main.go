package jsmops

import (
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
}

type LogLevel int

const (
	LogError LogLevel = iota
	LogInfo
	LogDebug
)

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
func Init(cloudID, apiToken, userName string, options *ClientOptions) *APIClient {
	client := &APIClient{
		cloudID:  cloudID,
		userName: userName,
		apiToken: apiToken,
	}
	if options != nil {
		client.logger = options.Logger
		client.logLevel = options.Level
	}
	if client.logger != nil {
		client.logger.Infof("JSM ops Client initializing..., authorization = %s", userName)
	}

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

	return client
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
