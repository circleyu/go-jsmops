package routingrules

type RoutingRule struct {
	ID          string                 `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
	Order       int                    `json:"order,omitempty"`
}

type ListRoutingRulesResult struct {
	Links        RoutingRulesResponseLink `json:"links"`
	RoutingRules []RoutingRule            `json:"values"`
	Count        int                      `json:"count"`
}

type RoutingRulesResponseLink struct {
	Next string `json:"next"`
}

