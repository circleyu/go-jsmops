package alert

import "time"

type Alert struct {
	Seen           bool        `json:"seen,omitempty"`
	Id             string      `json:"id,omitempty"`
	TinyID         string      `json:"tinyId,omitempty"`
	Alias          string      `json:"alias,omitempty"`
	Message        string      `json:"message,omitempty"`
	Status         string      `json:"status,omitempty"`
	Acknowledged   bool        `json:"acknowledged,omitempty"`
	IsSeen         bool        `json:"isSeen,omitempty"`
	Tags           []string    `json:"tags,omitempty"`
	Snoozed        bool        `json:"snoozed,omitempty"`
	SnoozedUntil   time.Time   `json:"snoozedUntil,omitempty"`
	Count          int         `json:"count,omitempty"`
	LastOccurredAt time.Time   `json:"lastOccurredAt,omitempty"`
	CreatedAt      time.Time   `json:"createdAt,omitempty"`
	UpdatedAt      time.Time   `json:"updatedAt,omitempty"`
	Source         string      `json:"source,omitempty"`
	Owner          string      `json:"owner,omitempty"`
	Priority       Priority    `json:"priority,omitempty"`
	Responders     []Responder `json:"responders"`
	Integration    Integration `json:"integration,omitempty"`
	Report         Report      `json:"report,omitempty"`
	OwnerTeamID    string      `json:"ownerTeamId,omitempty"`
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
	Result    string `json:"result"`
	RequestID string `json:"requestId"`
	Took      int    `json:"took"`
}

type AddNoteResponse struct {
	ID        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Note      string    `json:"note,omitempty"`
	Owner     string    `json:"owner,omitempty"`
}
