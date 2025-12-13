package filters

type GetIntegrationAlertFilterRequest struct {
	IntegrationID string
}

type UpdateIntegrationAlertFilterRequest struct {
	IntegrationID string
	Filter        map[string]interface{} `json:"filter"`
}

