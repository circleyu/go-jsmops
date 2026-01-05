package alert

import "time"

type Alert struct {
	Seen            bool                   `json:"seen,omitempty"`
	Id              string                 `json:"id,omitempty"`
	TinyID          string                 `json:"tinyId"`
	Alias           string                 `json:"alias"`
	Message         string                 `json:"message"`
	Status          string                 `json:"status"`
	Acknowledged    bool                   `json:"acknowledged,omitempty"`
	IsSeen          bool                   `json:"isSeen,omitempty"`
	Tags            []string               `json:"tags"`
	Snoozed         bool                   `json:"snoozed,omitempty"`
	SnoozedUntil    time.Time              `json:"snoozedUntil,omitempty"`
	Count           int                    `json:"count,omitempty"`
	LastOccurredAt  time.Time              `json:"lastOccurredAt,omitempty"`
	CreatedAt       time.Time              `json:"createdAt,omitempty"`
	UpdatedAt       time.Time              `json:"updatedAt,omitempty"`
	Source          string                 `json:"source,omitempty"`
	Owner           string                 `json:"owner"`
	Priority        Priority               `json:"priority"`
	Responders      []Responder            `json:"responders"`
	Integration     Integration            `json:"integration,omitempty"`
	Report          Report                 `json:"report,omitempty"`
	OwnerTeamID     string                 `json:"ownerTeamId,omitempty"`
	Entity          string                 `json:"entity,omitempty"`
	Description     string                 `json:"description"`
	Actions         []string               `json:"actions,omitempty"`
	ExtraProperties map[string]interface{} `json:"extraProperties"`
	IntegrationType string                 `json:"integrationType,omitempty"`
	IntegrationName string                 `json:"integrationName,omitempty"`
}

type Integration struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type ListAlertsResult struct {
	Links  ListAlertResponseLink `json:"links"`
	Alerts []Alert               `json:"values"`
	Count  int                   `json:"count"`
}

type ListAlertResponseLink struct {
	Next string `json:"next"`
}

type SuccessResponse struct {
	Result    string  `json:"result"`
	RequestID string  `json:"requestId"`
	Took      float32 `json:"took"`
}

type AddNoteResponse struct {
	ID        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Note      string    `json:"note,omitempty"`
	Owner     string    `json:"owner,omitempty"`
}

type AlertNote struct {
	ID        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Note      string    `json:"note,omitempty"`
	Owner     string    `json:"owner,omitempty"`
}

type ListAlertNotesResult struct {
	Links ListAlertResponseLink `json:"links"`
	Notes []AlertNote           `json:"values"`
	Count int                   `json:"count"`
}

type AlertLog struct {
	Log       string    `json:"log,omitempty"`
	LogType   string    `json:"logType,omitempty"`
	Owner     string    `json:"owner,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type ListAlertLogsResult struct {
	Links ListAlertResponseLink `json:"links"`
	Logs  []AlertLog            `json:"values"`
	Count int                   `json:"count"`
}

type RequestStatusResponse struct {
	Status    string `json:"status,omitempty"`
	RequestID string `json:"requestId,omitempty"`
	Alert     *Alert `json:"alert,omitempty"`
}
