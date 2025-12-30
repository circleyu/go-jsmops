package policies

import (
	"strconv"

	"github.com/circleyu/go-jsmops/v2/params"
)

type ListTeamPoliciesRequest struct {
	TeamID string
	Type   string
	Offset int
	Size   int
}

func (r *ListTeamPoliciesRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Type != "" {
		query.Is("type", r.Type)
	}
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type CreateTeamPolicyRequest struct {
	TeamID      string
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
	Order       int                    `json:"order,omitempty"`
}

type GetTeamPolicyRequest struct {
	TeamID   string
	PolicyID string
}

type PutTeamPolicyRequest struct {
	TeamID      string
	PolicyID    string
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
	Order       int                    `json:"order,omitempty"`
}

type DeleteTeamPolicyRequest struct {
	TeamID   string
	PolicyID string
}

type ChangeOrderTeamPolicyRequest struct {
	TeamID   string
	PolicyID string
	Order    int `json:"order"`
}

type EnableTeamPolicyRequest struct {
	TeamID   string
	PolicyID string
}

type DisableTeamPolicyRequest struct {
	TeamID   string
	PolicyID string
}

