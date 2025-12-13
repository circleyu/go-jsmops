package jsmops

type alertsEndpoints struct {
	GetRequestStatus        func(string) string
	ListAlerts              string
	CreateAlert             string
	GetAlert                func(string) string
	DeleteAlert             func(string) string
	GetAlertByAlias         string
	AcknowledgeAlert        func(string) string
	AssignAlert             func(string) string
	AddResponder            func(string) string
	AddExtraProperties      func(string) string
	DeleteExtraProperties   func(string) string
	AddTags                 func(string) string
	DeleteTags              func(string) string
	CloseAlert              func(string) string
	EscalateAlert           func(string) string
	ExecuteCustomAction     func(string) string
	SnoozeAlert             func(string) string
	UnacknowledgeAlert      func(string) string
	ListAlertNotes          func(string) string
	AddAlertNote            func(string) string
	DeleteAlertNote         func(string, string) string
	UpdateAlertNote         func(string, string) string
	UpdateAlertPriority     func(string) string
	UpdateAlertMessage      func(string) string
	UpdateAlertDescription  func(string) string
	ListAlertLogs           func(string) string
}

type auditLogsEndpoints struct {
	GetAuditLogs string
}

type contactsEndpoints struct {
	ListContacts     string
	CreateContact    string
	GetContact       func(string) string
	DeleteContact    func(string) string
	UpdateContact    func(string) string
	ActivateContact  func(string) string
	DeactivateContact func(string) string
}

type teamsEndpoints struct {
	ListTeams           string
	EnableOps           func(string) string
	GetTeamRequestStatus func(string, string) string
}

type rolesEndpoints struct {
	ListCustomUserRoles string
	GetCustomUserRole   func(string) string
	CreateCustomUserRole string
	UpdateCustomUserRole func(string) string
	DeleteCustomUserRole func(string) string
	AssignCustomUserRole string
}

type escalationsEndpoints struct {
	ListEscalations   func(string) string
	CreateEscalation  func(string) string
	GetEscalation     func(string, string) string
	UpdateEscalation  func(string, string) string
	DeleteEscalation  func(string, string) string
}

type forwardingRulesEndpoints struct {
	ListForwardingRules  string
	CreateForwardingRule string
	GetForwardingRule    func(string) string
	UpdateForwardingRule func(string) string
	DeleteForwardingRule func(string) string
}

type heartbeatsEndpoints struct {
	ListHeartbeats  func(string) string
	CreateHeartbeat func(string) string
	UpdateHeartbeat func(string) string
	DeleteHeartbeat func(string) string
	PingHeartbeat   func(string) string
}

type integrationsEndpoints struct {
	ListIntegrations  string
	CreateIntegration string
	GetIntegration    func(string) string
	UpdateIntegration func(string) string
	DeleteIntegration func(string) string
}

type integrationActionsEndpoints struct {
	ListIntegrationActions   func(string) string
	CreateIntegrationAction   func(string) string
	GetIntegrationAction      func(string, string) string
	UpdateIntegrationAction   func(string, string) string
	DeleteIntegrationAction   func(string, string) string
	ReorderIntegrationAction  func(string, string) string
}

type integrationFiltersEndpoints struct {
	GetIntegrationAlertFilter    func(string) string
	UpdateIntegrationAlertFilter func(string) string
}

type maintenancesEndpoints struct {
	ListGlobalMaintenances  string
	CreateGlobalMaintenance string
	GetGlobalMaintenance    func(string) string
	UpdateGlobalMaintenance func(string) string
	DeleteGlobalMaintenance func(string) string
	CancelGlobalMaintenance func(string) string
	ListTeamMaintenances    func(string) string
	CreateTeamMaintenance   func(string) string
	GetTeamMaintenance      func(string, string) string
	UpdateTeamMaintenance   func(string, string) string
	DeleteTeamMaintenance   func(string, string) string
	CancelTeamMaintenance   func(string, string) string
}

type notificationRulesEndpoints struct {
	ListNotificationRules string
	CreateNotificationRule string
	GetNotificationRule   func(string) string
	UpdateNotificationRule func(string) string
	DeleteNotificationRule func(string) string
}

type notificationRuleStepsEndpoints struct {
	ListNotificationRuleSteps func(string) string
	CreateNotificationRuleStep func(string) string
	GetNotificationRuleStep   func(string, string) string
	UpdateNotificationRuleStep func(string, string) string
	DeleteNotificationRuleStep func(string, string) string
}

type policiesEndpoints struct {
	ListGlobalAlertPolicies string
	CreateGlobalAlertPolicy string
	GetGlobalAlertPolicy   func(string) string
	PutGlobalAlertPolicy   func(string) string
	DeleteGlobalAlertPolicy func(string) string
	ChangeOrderGlobalAlertPolicy func(string) string
	EnableGlobalAlertPolicy func(string) string
	DisableGlobalAlertPolicy func(string) string
}

type teamPoliciesEndpoints struct {
	ListTeamPolicies func(string) string
	CreateTeamPolicy func(string) string
	GetTeamPolicy    func(string, string) string
	PutTeamPolicy    func(string, string) string
	DeleteTeamPolicy func(string, string) string
	ChangeOrderTeamPolicy func(string, string) string
	EnableTeamPolicy func(string, string) string
	DisableTeamPolicy func(string, string) string
}

type teamRolesEndpoints struct {
	ListTeamRoles func(string) string
	GetTeamRole   func(string, string) string
	CreateTeamRole func(string) string
	UpdateTeamRole func(string, string) string
	DeleteTeamRole func(string, string) string
}

type routingRulesEndpoints struct {
	ListRoutingRules func(string) string
	CreateRoutingRule func(string) string
	GetRoutingRule   func(string, string) string
	UpdateRoutingRule func(string, string) string
	DeleteRoutingRule func(string, string) string
	ChangeOrderRoutingRule func(string, string) string
}

type schedulesEndpoints struct {
	ListSchedules string
	CreateSchedule string
	GetSchedule   func(string) string
	UpdateSchedule func(string) string
	DeleteSchedule func(string) string
}

type schedulesOnCallsEndpoints struct {
	ListOnCallResponders func(string) string
	ListNextOnCallResponders func(string) string
	ExportOnCallResponders func(string) string
}

type schedulesOverridesEndpoints struct {
	ListOverrides func(string) string
	CreateOverride func(string) string
	GetOverride   func(string, string) string
	UpdateOverride func(string, string) string
	DeleteOverride func(string, string) string
}

type schedulesRotationsEndpoints struct {
	ListRotations func(string) string
	CreateRotation func(string) string
	GetRotation   func(string, string) string
	UpdateRotation func(string, string) string
	DeleteRotation func(string, string) string
}

type schedulesTimelinesEndpoints struct {
	GetScheduleTimeline func(string) string
	ExportScheduleTimeline func(string) string
}

type syncsEndpoints struct {
	ListSyncs string
	CreateSync string
	GetSync   func(string) string
	UpdateSync func(string) string
	DeleteSync func(string) string
}

type syncsActionsEndpoints struct {
	ListSyncActions func(string) string
	CreateSyncAction func(string) string
	GetSyncAction   func(string, string) string
	UpdateSyncAction func(string, string) string
	DeleteSyncAction func(string, string) string
	ReorderSyncAction func(string, string) string
}

type syncsActionGroupsEndpoints struct {
	ListSyncActionGroups func(string) string
	CreateSyncActionGroup func(string) string
	GetSyncActionGroup   func(string, string) string
	UpdateSyncActionGroup func(string, string) string
	DeleteSyncActionGroup func(string, string) string
	ReorderSyncActionGroup func(string, string) string
}

type jecEndpoints struct {
	ListJECChannels string
	CreateJECChannel string
	GetJECChannel   func(string) string
	DeleteJECChannel func(string) string
	SendJECAction   string
}

var endpoints = struct {
	alerts             alertsEndpoints
	auditLogs          auditLogsEndpoints
	contacts           contactsEndpoints
	teams              teamsEndpoints
	roles              rolesEndpoints
	escalations        escalationsEndpoints
	forwardingRules    forwardingRulesEndpoints
	heartbeats         heartbeatsEndpoints
	integrations       integrationsEndpoints
	integrationActions integrationActionsEndpoints
	integrationFilters integrationFiltersEndpoints
	maintenances       maintenancesEndpoints
	notificationRules  notificationRulesEndpoints
	notificationRuleSteps notificationRuleStepsEndpoints
	policies           policiesEndpoints
	teamPolicies       teamPoliciesEndpoints
	teamRoles          teamRolesEndpoints
	routingRules       routingRulesEndpoints
	schedules          schedulesEndpoints
	schedulesOnCalls   schedulesOnCallsEndpoints
	schedulesOverrides schedulesOverridesEndpoints
	schedulesRotations schedulesRotationsEndpoints
	schedulesTimelines schedulesTimelinesEndpoints
	syncs              syncsEndpoints
	syncsActions       syncsActionsEndpoints
	syncsActionGroups  syncsActionGroupsEndpoints
	jec                jecEndpoints
}{
	alerts: alertsEndpoints{
		GetRequestStatus:       func(id string) string { return "v1/alerts/requests/" + id },
		ListAlerts:             "v1/alerts",
		CreateAlert:            "v1/alerts",
		GetAlert:               func(id string) string { return "v1/alerts/" + id },
		DeleteAlert:            func(id string) string { return "v1/alerts/" + id },
		GetAlertByAlias:        "v1/alerts/alias",
		AcknowledgeAlert:        func(id string) string { return "v1/alerts/" + id + "/acknowledge" },
		AssignAlert:            func(id string) string { return "v1/alerts/" + id + "/assign" },
		AddResponder:           func(id string) string { return "v1/alerts/" + id + "/responders" },
		AddExtraProperties:     func(id string) string { return "v1/alerts/" + id + "/extra-properties" },
		DeleteExtraProperties:  func(id string) string { return "v1/alerts/" + id + "/extra-properties" },
		AddTags:                func(id string) string { return "v1/alerts/" + id + "/tags" },
		DeleteTags:             func(id string) string { return "v1/alerts/" + id + "/tags" },
		CloseAlert:             func(id string) string { return "v1/alerts/" + id + "/close" },
		EscalateAlert:          func(id string) string { return "v1/alerts/" + id + "/escalate" },
		ExecuteCustomAction:    func(id string) string { return "v1/alerts/" + id + "/action" },
		SnoozeAlert:            func(id string) string { return "v1/alerts/" + id + "/snooze" },
		UnacknowledgeAlert:     func(id string) string { return "v1/alerts/" + id + "/unacknowledge" },
		ListAlertNotes:         func(id string) string { return "v1/alerts/" + id + "/notes" },
		AddAlertNote:           func(id string) string { return "v1/alerts/" + id + "/notes" },
		DeleteAlertNote:        func(alertId, noteId string) string { return "v1/alerts/" + alertId + "/notes/" + noteId },
		UpdateAlertNote:        func(alertId, noteId string) string { return "v1/alerts/" + alertId + "/notes/" + noteId },
		UpdateAlertPriority:    func(id string) string { return "v1/alerts/" + id + "/priority" },
		UpdateAlertMessage:     func(id string) string { return "v1/alerts/" + id + "/message" },
		UpdateAlertDescription: func(id string) string { return "v1/alerts/" + id + "/description" },
		ListAlertLogs:          func(id string) string { return "v1/alerts/" + id + "/logs" },
	},
	auditLogs: auditLogsEndpoints{
		GetAuditLogs: "v1/logs",
	},
	contacts: contactsEndpoints{
		ListContacts:     "v1/users/contacts",
		CreateContact:    "v1/users/contacts",
		GetContact:       func(id string) string { return "v1/users/contacts/" + id },
		DeleteContact:    func(id string) string { return "v1/users/contacts/" + id },
		UpdateContact:    func(id string) string { return "v1/users/contacts/" + id },
		ActivateContact:  func(id string) string { return "v1/users/contacts/" + id + "/activate" },
		DeactivateContact: func(id string) string { return "v1/users/contacts/" + id + "/deactivate" },
	},
	teams: teamsEndpoints{
		ListTeams:           "v1/teams",
		EnableOps:           func(teamId string) string { return "v1/teams/" + teamId + "/enable-ops" },
		GetTeamRequestStatus: func(teamId, requestId string) string { return "v1/teams/" + teamId + "/requests/" + requestId },
	},
	roles: rolesEndpoints{
		ListCustomUserRoles: "v1/roles",
		GetCustomUserRole:   func(identifier string) string { return "v1/roles/" + identifier },
		CreateCustomUserRole: "v1/roles",
		UpdateCustomUserRole: func(identifier string) string { return "v1/roles/" + identifier },
		DeleteCustomUserRole: func(identifier string) string { return "v1/roles/" + identifier },
		AssignCustomUserRole: "v1/roles/assign",
	},
	escalations: escalationsEndpoints{
		ListEscalations:  func(teamId string) string { return "v1/teams/" + teamId + "/escalations" },
		CreateEscalation: func(teamId string) string { return "v1/teams/" + teamId + "/escalations" },
		GetEscalation:    func(teamId, id string) string { return "v1/teams/" + teamId + "/escalations/" + id },
		UpdateEscalation: func(teamId, id string) string { return "v1/teams/" + teamId + "/escalations/" + id },
		DeleteEscalation: func(teamId, id string) string { return "v1/teams/" + teamId + "/escalations/" + id },
	},
	forwardingRules: forwardingRulesEndpoints{
		ListForwardingRules:  "v1/forwarding-rules",
		CreateForwardingRule: "v1/forwarding-rules",
		GetForwardingRule:    func(id string) string { return "v1/forwarding-rules/" + id },
		UpdateForwardingRule: func(id string) string { return "v1/forwarding-rules/" + id },
		DeleteForwardingRule: func(id string) string { return "v1/forwarding-rules/" + id },
	},
	heartbeats: heartbeatsEndpoints{
		ListHeartbeats:  func(teamId string) string { return "v1/teams/" + teamId + "/heartbeats" },
		CreateHeartbeat: func(teamId string) string { return "v1/teams/" + teamId + "/heartbeats" },
		UpdateHeartbeat: func(teamId string) string { return "v1/teams/" + teamId + "/heartbeats" },
		DeleteHeartbeat: func(teamId string) string { return "v1/teams/" + teamId + "/heartbeats" },
		PingHeartbeat:   func(teamId string) string { return "v1/teams/" + teamId + "/heartbeats/ping" },
	},
	integrations: integrationsEndpoints{
		ListIntegrations:  "v1/integrations",
		CreateIntegration: "v1/integrations",
		GetIntegration:    func(id string) string { return "v1/integrations/" + id },
		UpdateIntegration: func(id string) string { return "v1/integrations/" + id },
		DeleteIntegration: func(id string) string { return "v1/integrations/" + id },
	},
	integrationActions: integrationActionsEndpoints{
		ListIntegrationActions:  func(integrationId string) string { return "v1/integrations/" + integrationId + "/actions" },
		CreateIntegrationAction:  func(integrationId string) string { return "v1/integrations/" + integrationId + "/actions" },
		GetIntegrationAction:    func(integrationId, id string) string { return "v1/integrations/" + integrationId + "/actions/" + id },
		UpdateIntegrationAction:  func(integrationId, id string) string { return "v1/integrations/" + integrationId + "/actions/" + id },
		DeleteIntegrationAction:  func(integrationId, id string) string { return "v1/integrations/" + integrationId + "/actions/" + id },
		ReorderIntegrationAction: func(integrationId, id string) string { return "v1/integrations/" + integrationId + "/actions/" + id + "/order" },
	},
	integrationFilters: integrationFiltersEndpoints{
		GetIntegrationAlertFilter:    func(integrationId string) string { return "v1/integrations/" + integrationId + "/outgoing/alert-filter/main" },
		UpdateIntegrationAlertFilter: func(integrationId string) string { return "v1/integrations/" + integrationId + "/outgoing/alert-filter/main" },
	},
	maintenances: maintenancesEndpoints{
		ListGlobalMaintenances:  "v1/maintenances",
		CreateGlobalMaintenance: "v1/maintenances",
		GetGlobalMaintenance:    func(id string) string { return "v1/maintenances/" + id },
		UpdateGlobalMaintenance: func(id string) string { return "v1/maintenances/" + id },
		DeleteGlobalMaintenance: func(id string) string { return "v1/maintenances/" + id },
		CancelGlobalMaintenance: func(id string) string { return "v1/maintenances/" + id + "/cancel" },
		ListTeamMaintenances:    func(teamId string) string { return "v1/teams/" + teamId + "/maintenances" },
		CreateTeamMaintenance:   func(teamId string) string { return "v1/teams/" + teamId + "/maintenances" },
		GetTeamMaintenance:      func(teamId, id string) string { return "v1/teams/" + teamId + "/maintenances/" + id },
		UpdateTeamMaintenance:   func(teamId, id string) string { return "v1/teams/" + teamId + "/maintenances/" + id },
		DeleteTeamMaintenance:   func(teamId, id string) string { return "v1/teams/" + teamId + "/maintenances/" + id },
		CancelTeamMaintenance:   func(teamId, id string) string { return "v1/teams/" + teamId + "/maintenances/" + id + "/cancel" },
	},
	notificationRules: notificationRulesEndpoints{
		ListNotificationRules: "v1/notification-rules",
		CreateNotificationRule: "v1/notification-rules",
		GetNotificationRule:   func(id string) string { return "v1/notification-rules/" + id },
		UpdateNotificationRule: func(id string) string { return "v1/notification-rules/" + id },
		DeleteNotificationRule: func(id string) string { return "v1/notification-rules/" + id },
	},
	notificationRuleSteps: notificationRuleStepsEndpoints{
		ListNotificationRuleSteps: func(ruleId string) string { return "v1/notification-rules/" + ruleId + "/steps" },
		CreateNotificationRuleStep: func(ruleId string) string { return "v1/notification-rules/" + ruleId + "/steps" },
		GetNotificationRuleStep:   func(ruleId, id string) string { return "v1/notification-rules/" + ruleId + "/steps/" + id },
		UpdateNotificationRuleStep: func(ruleId, id string) string { return "v1/notification-rules/" + ruleId + "/steps/" + id },
		DeleteNotificationRuleStep: func(ruleId, id string) string { return "v1/notification-rules/" + ruleId + "/steps/" + id },
	},
	policies: policiesEndpoints{
		ListGlobalAlertPolicies: "v1/alerts/policies",
		CreateGlobalAlertPolicy: "v1/alerts/policies",
		GetGlobalAlertPolicy:   func(policyId string) string { return "v1/alerts/policies/" + policyId },
		PutGlobalAlertPolicy:   func(policyId string) string { return "v1/alerts/policies/" + policyId },
		DeleteGlobalAlertPolicy: func(policyId string) string { return "v1/alerts/policies/" + policyId },
		ChangeOrderGlobalAlertPolicy: func(policyId string) string { return "v1/alerts/policies/" + policyId + "/change-order" },
		EnableGlobalAlertPolicy: func(policyId string) string { return "v1/alerts/policies/" + policyId + "/enable" },
		DisableGlobalAlertPolicy: func(policyId string) string { return "v1/alerts/policies/" + policyId + "/disable" },
	},
	teamPolicies: teamPoliciesEndpoints{
		ListTeamPolicies: func(teamId string) string { return "v1/teams/" + teamId + "/policies" },
		CreateTeamPolicy: func(teamId string) string { return "v1/teams/" + teamId + "/policies" },
		GetTeamPolicy:    func(teamId, policyId string) string { return "v1/teams/" + teamId + "/policies/" + policyId },
		PutTeamPolicy:    func(teamId, policyId string) string { return "v1/teams/" + teamId + "/policies/" + policyId },
		DeleteTeamPolicy: func(teamId, policyId string) string { return "v1/teams/" + teamId + "/policies/" + policyId },
		ChangeOrderTeamPolicy: func(teamId, policyId string) string { return "v1/teams/" + teamId + "/policies/" + policyId + "/change-order" },
		EnableTeamPolicy: func(teamId, policyId string) string { return "v1/teams/" + teamId + "/policies/" + policyId + "/enable" },
		DisableTeamPolicy: func(teamId, policyId string) string { return "v1/teams/" + teamId + "/policies/" + policyId + "/disable" },
	},
	teamRoles: teamRolesEndpoints{
		ListTeamRoles: func(teamId string) string { return "v1/teams/" + teamId + "/roles" },
		GetTeamRole:   func(teamId, identifier string) string { return "v1/teams/" + teamId + "/roles/" + identifier },
		CreateTeamRole: func(teamId string) string { return "v1/teams/" + teamId + "/roles" },
		UpdateTeamRole: func(teamId, identifier string) string { return "v1/teams/" + teamId + "/roles/" + identifier },
		DeleteTeamRole: func(teamId, identifier string) string { return "v1/teams/" + teamId + "/roles/" + identifier },
	},
	routingRules: routingRulesEndpoints{
		ListRoutingRules: func(teamId string) string { return "v1/teams/" + teamId + "/routing-rules" },
		CreateRoutingRule: func(teamId string) string { return "v1/teams/" + teamId + "/routing-rules" },
		GetRoutingRule:   func(teamId, id string) string { return "v1/teams/" + teamId + "/routing-rules/" + id },
		UpdateRoutingRule: func(teamId, id string) string { return "v1/teams/" + teamId + "/routing-rules/" + id },
		DeleteRoutingRule: func(teamId, id string) string { return "v1/teams/" + teamId + "/routing-rules/" + id },
		ChangeOrderRoutingRule: func(teamId, id string) string { return "v1/teams/" + teamId + "/routing-rules/" + id + "/change-order" },
	},
	schedules: schedulesEndpoints{
		ListSchedules: "v1/schedules",
		CreateSchedule: "v1/schedules",
		GetSchedule:   func(id string) string { return "v1/schedules/" + id },
		UpdateSchedule: func(id string) string { return "v1/schedules/" + id },
		DeleteSchedule: func(id string) string { return "v1/schedules/" + id },
	},
	schedulesOnCalls: schedulesOnCallsEndpoints{
		ListOnCallResponders: func(scheduleId string) string { return "v1/schedules/" + scheduleId + "/on-calls" },
		ListNextOnCallResponders: func(scheduleId string) string { return "v1/schedules/" + scheduleId + "/next-on-calls" },
		ExportOnCallResponders: func(userIdentifier string) string { return "v1/schedules/on-calls/" + userIdentifier + ".ics" },
	},
	schedulesOverrides: schedulesOverridesEndpoints{
		ListOverrides: func(scheduleId string) string { return "v1/schedules/" + scheduleId + "/overrides" },
		CreateOverride: func(scheduleId string) string { return "v1/schedules/" + scheduleId + "/overrides" },
		GetOverride:   func(scheduleId, alias string) string { return "v1/schedules/" + scheduleId + "/overrides/" + alias },
		UpdateOverride: func(scheduleId, alias string) string { return "v1/schedules/" + scheduleId + "/overrides/" + alias },
		DeleteOverride: func(scheduleId, alias string) string { return "v1/schedules/" + scheduleId + "/overrides/" + alias },
	},
	schedulesRotations: schedulesRotationsEndpoints{
		ListRotations: func(scheduleId string) string { return "v1/schedules/" + scheduleId + "/rotations" },
		CreateRotation: func(scheduleId string) string { return "v1/schedules/" + scheduleId + "/rotations" },
		GetRotation:   func(scheduleId, id string) string { return "v1/schedules/" + scheduleId + "/rotations/" + id },
		UpdateRotation: func(scheduleId, id string) string { return "v1/schedules/" + scheduleId + "/rotations/" + id },
		DeleteRotation: func(scheduleId, id string) string { return "v1/schedules/" + scheduleId + "/rotations/" + id },
	},
	schedulesTimelines: schedulesTimelinesEndpoints{
		GetScheduleTimeline: func(scheduleId string) string { return "v1/schedules/" + scheduleId + "/timeline" },
		ExportScheduleTimeline: func(scheduleId string) string { return "v1/schedules/" + scheduleId + ".ics" },
	},
	syncs: syncsEndpoints{
		ListSyncs: "v1/syncs",
		CreateSync: "v1/syncs",
		GetSync:   func(id string) string { return "v1/syncs/" + id },
		UpdateSync: func(id string) string { return "v1/syncs/" + id },
		DeleteSync: func(id string) string { return "v1/syncs/" + id },
	},
	syncsActions: syncsActionsEndpoints{
		ListSyncActions: func(syncId string) string { return "v1/syncs/" + syncId + "/actions" },
		CreateSyncAction: func(syncId string) string { return "v1/syncs/" + syncId + "/actions" },
		GetSyncAction:   func(syncId, id string) string { return "v1/syncs/" + syncId + "/actions/" + id },
		UpdateSyncAction: func(syncId, id string) string { return "v1/syncs/" + syncId + "/actions/" + id },
		DeleteSyncAction: func(syncId, id string) string { return "v1/syncs/" + syncId + "/actions/" + id },
		ReorderSyncAction: func(syncId, id string) string { return "v1/syncs/" + syncId + "/actions/" + id + "/order" },
	},
	syncsActionGroups: syncsActionGroupsEndpoints{
		ListSyncActionGroups: func(syncId string) string { return "v1/syncs/" + syncId + "/action-groups" },
		CreateSyncActionGroup: func(syncId string) string { return "v1/syncs/" + syncId + "/action-groups" },
		GetSyncActionGroup:   func(syncId, id string) string { return "v1/syncs/" + syncId + "/action-groups/" + id },
		UpdateSyncActionGroup: func(syncId, id string) string { return "v1/syncs/" + syncId + "/action-groups/" + id },
		DeleteSyncActionGroup: func(syncId, id string) string { return "v1/syncs/" + syncId + "/action-groups/" + id },
		ReorderSyncActionGroup: func(syncId, id string) string { return "v1/syncs/" + syncId + "/action-groups/" + id + "/order" },
	},
	jec: jecEndpoints{
		ListJECChannels: "v1/jec/channels",
		CreateJECChannel: "v1/jec/channels",
		GetJECChannel:   func(id string) string { return "v1/jec/channels/" + id },
		DeleteJECChannel: func(id string) string { return "v1/jec/channels/" + id },
		SendJECAction:   "v1/jec/action",
	},
}
