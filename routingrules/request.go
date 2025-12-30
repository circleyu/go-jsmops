package routingrules

import (
	"strconv"

	"github.com/circleyu/go-jsmops/v2/params"
)

type ListRoutingRulesRequest struct {
	TeamID string
	Offset int
	Size   int
}

func (r *ListRoutingRulesRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type CreateRoutingRuleRequest struct {
	TeamID      string
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
	Order       int                    `json:"order,omitempty"`
}

type GetRoutingRuleRequest struct {
	TeamID string
	ID     string
}

type UpdateRoutingRuleRequest struct {
	TeamID      string
	ID          string
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
	Order       int                    `json:"order,omitempty"`
}

type DeleteRoutingRuleRequest struct {
	TeamID string
	ID     string
}

type ChangeOrderRoutingRuleRequest struct {
	TeamID string
	ID     string
	Order  int `json:"order"`
}

