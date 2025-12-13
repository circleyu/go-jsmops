package escalations

import (
	"strconv"

	"github.com/circleyu/go-jsmops/params"
)

type ListEscalationsRequest struct {
	TeamID string
	Offset int
	Size   int
}

func (r *ListEscalationsRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type CreateEscalationRequest struct {
	TeamID      string
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Rules       []EscalationRule       `json:"rules,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Repeat      *EscalationRepeat      `json:"repeat,omitempty"`
}

type EscalationRule struct {
	Condition  string                 `json:"condition,omitempty"`
	NotifyType string                 `json:"notifyType,omitempty"`
	Delay      int                    `json:"delay,omitempty"`
	Recipient  map[string]interface{} `json:"recipient,omitempty"`
}

type EscalationRepeat struct {
	WaitInterval        int  `json:"waitInterval,omitempty"`
	Count               int  `json:"count,omitempty"`
	ResetRecipientStates bool `json:"resetRecipientStates,omitempty"`
	CloseAlertAfterAll  bool `json:"closeAlertAfterAll,omitempty"`
}

type GetEscalationRequest struct {
	TeamID string
	ID     string
}

type UpdateEscalationRequest struct {
	TeamID      string
	ID          string
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Rules       []EscalationRule  `json:"rules,omitempty"`
	Enabled     *bool             `json:"enabled,omitempty"`
	Repeat      *EscalationRepeat `json:"repeat,omitempty"`
}

type DeleteEscalationRequest struct {
	TeamID string
	ID     string
}

