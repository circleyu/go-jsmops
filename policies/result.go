package policies

type Policy struct {
	ID          string                 `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
	Order       int                    `json:"order,omitempty"`
}

type ListPoliciesResult struct {
	Links    PoliciesResponseLink `json:"links"`
	Policies []Policy             `json:"values"`
	Count    int                  `json:"count"`
}

type PoliciesResponseLink struct {
	Next string `json:"next"`
}

