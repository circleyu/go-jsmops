package actiongroups

type SyncActionGroup struct {
	ID      string                 `json:"id,omitempty"`
	Name    string                 `json:"name,omitempty"`
	Order   int                    `json:"order,omitempty"`
	Enabled bool                   `json:"enabled,omitempty"`
	Config  map[string]interface{} `json:"config,omitempty"`
}

type ListSyncActionGroupsResult struct {
	Links       SyncActionGroupsResponseLink `json:"links"`
	ActionGroups []SyncActionGroup            `json:"values"`
	Count        int                          `json:"count"`
}

type SyncActionGroupsResponseLink struct {
	Next string `json:"next"`
}

