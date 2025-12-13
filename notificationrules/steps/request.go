package steps

import (
	"strconv"

	"github.com/circleyu/go-jsmops/params"
)

type ListNotificationRuleStepsRequest struct {
	RuleID string
	Offset int
	Size   int
}

func (r *ListNotificationRuleStepsRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type CreateNotificationRuleStepRequest struct {
	RuleID      string
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
	Order       int                    `json:"order,omitempty"`
}

type GetNotificationRuleStepRequest struct {
	RuleID string
	ID     string
}

type UpdateNotificationRuleStepRequest struct {
	RuleID      string
	ID          string
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
	Order       int                    `json:"order,omitempty"`
}

type DeleteNotificationRuleStepRequest struct {
	RuleID string
	ID     string
}

