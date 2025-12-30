package policies

import (
	"strconv"

	"github.com/circleyu/go-jsmops/v2/params"
)

type ListGlobalAlertPoliciesRequest struct {
	Offset int
	Size   int
}

func (r *ListGlobalAlertPoliciesRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type CreateGlobalAlertPolicyRequest struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
	Order       int                    `json:"order,omitempty"`
}

type GetGlobalAlertPolicyRequest struct {
	PolicyID string
}

type PutGlobalAlertPolicyRequest struct {
	PolicyID    string
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
	Order       int                    `json:"order,omitempty"`
}

type DeleteGlobalAlertPolicyRequest struct {
	PolicyID string
}

type ChangeOrderGlobalAlertPolicyRequest struct {
	PolicyID string
	Order    int `json:"order"`
}

type EnableGlobalAlertPolicyRequest struct {
	PolicyID string
}

type DisableGlobalAlertPolicyRequest struct {
	PolicyID string
}

