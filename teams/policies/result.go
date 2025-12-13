package policies

type TeamPolicy struct {
	ID          string                 `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
	Order       int                    `json:"order,omitempty"`
}

type ListTeamPoliciesResult struct {
	Links    TeamPoliciesResponseLink `json:"links"`
	Policies []TeamPolicy              `json:"values"`
	Count    int                       `json:"count"`
}

type TeamPoliciesResponseLink struct {
	Next string `json:"next"`
}

