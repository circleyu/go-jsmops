package jsmops

type alertsEndpoints struct {
	ListAlerts       string
	CreateAlert      string
	AcknowledgeAlert func(string) string
	CloseAlert       func(string) string
	AddAlertNote     func(string) string
}

var endpoints = struct {
	alerts alertsEndpoints
}{
	alerts: alertsEndpoints{
		ListAlerts:       "v1/alerts",
		CreateAlert:      "v1/alerts",
		AcknowledgeAlert: func(id string) string { return "v1/alerts/" + id + "/acknowledge" },
		CloseAlert:       func(id string) string { return "v1/alerts/" + id + "/close" },
		AddAlertNote:     func(id string) string { return "v1/alerts/" + id + "/notes" },
	},
}
