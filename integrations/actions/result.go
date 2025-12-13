package actions

type IntegrationAction struct {
	ID   string                 `json:"id,omitempty"`
	Type string                 `json:"type,omitempty"`
	Name string                 `json:"name,omitempty"`
	Config map[string]interface{} `json:"config,omitempty"`
	Order int                    `json:"order,omitempty"`
}

type ListIntegrationActionsResult struct {
	Links   IntegrationActionsResponseLink `json:"links"`
	Actions []IntegrationAction             `json:"values"`
	Count   int                             `json:"count"`
}

type IntegrationActionsResponseLink struct {
	Next string `json:"next"`
}

