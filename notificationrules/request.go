package notificationrules

import (
	"strconv"

	"github.com/circleyu/go-jsmops/v2/params"
)

type ListNotificationRulesRequest struct {
	Offset int
	Size   int
}

func (r *ListNotificationRulesRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type CreateNotificationRuleRequest struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
}

type GetNotificationRuleRequest struct {
	ID string
}

type UpdateNotificationRuleRequest struct {
	ID          string
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
}

type DeleteNotificationRuleRequest struct {
	ID string
}

