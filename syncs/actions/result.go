package actions

type SyncAction struct {
	ID      string                 `json:"id,omitempty"`
	Name    string                 `json:"name,omitempty"`
	Type    string                 `json:"type,omitempty"`
	Order   int                    `json:"order,omitempty"`
	Enabled bool                   `json:"enabled,omitempty"`
	Config  map[string]interface{} `json:"config,omitempty"`
}

type ListSyncActionsResult struct {
	Links   SyncActionsResponseLink `json:"links"`
	Actions []SyncAction             `json:"values"`
	Count   int                      `json:"count"`
}

type SyncActionsResponseLink struct {
	Next string `json:"next"`
}

