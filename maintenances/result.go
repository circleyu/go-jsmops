package maintenances

type Maintenance struct {
	ID          string                 `json:"id,omitempty"`
	Description string                 `json:"description,omitempty"`
	StartTime   string                 `json:"startTime,omitempty"`
	EndTime     string                 `json:"endTime,omitempty"`
	Rules       map[string]interface{} `json:"rules,omitempty"`
	Status      string                 `json:"status,omitempty"`
}

type ListMaintenancesResult struct {
	Links        MaintenancesResponseLink `json:"links"`
	Maintenances []Maintenance            `json:"values"`
	Count        int                      `json:"count"`
}

type MaintenancesResponseLink struct {
	Next string `json:"next"`
}

