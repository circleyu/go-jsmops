package integrations

type Integration struct {
	ID     string                 `json:"id,omitempty"`
	Type   string                 `json:"type,omitempty"`
	Name   string                 `json:"name,omitempty"`
	TeamID string                 `json:"teamId,omitempty"`
	Config map[string]interface{} `json:"config,omitempty"`
}

type ListIntegrationsResult struct {
	Links        IntegrationsResponseLink `json:"links"`
	Integrations []Integration            `json:"values"`
	Count        int                      `json:"count"`
}

type IntegrationsResponseLink struct {
	Next string `json:"next"`
}

