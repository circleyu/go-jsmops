package schedules

type Schedule struct {
	ID          string                 `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	TeamID      string                 `json:"teamId,omitempty"`
	Timezone    string                 `json:"timezone,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
}

type ListSchedulesResult struct {
	Links     SchedulesResponseLink `json:"links"`
	Schedules []Schedule            `json:"values"`
	Count     int                   `json:"count"`
}

type SchedulesResponseLink struct {
	Next string `json:"next"`
}

