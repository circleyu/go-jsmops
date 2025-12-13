package syncs

import "time"

type Sync struct {
	ID          string                 `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Type        string                 `json:"type,omitempty"`
	TeamID      string                 `json:"teamId,omitempty"`
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
	CreatedAt   time.Time              `json:"createdAt,omitempty"`
	UpdatedAt   time.Time              `json:"updatedAt,omitempty"`
}

type ListSyncsResult struct {
	Links SyncsResponseLink `json:"links"`
	Syncs []Sync            `json:"values"`
	Count int               `json:"count"`
}

type SyncsResponseLink struct {
	Next string `json:"next"`
}

