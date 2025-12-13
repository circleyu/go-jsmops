package integrations

import (
	"strconv"

	"github.com/circleyu/go-jsmops/params"
)

type ListIntegrationsRequest struct {
	Type   string
	TeamID string
	Name   string
	Offset int
	Size   int
}

func (r *ListIntegrationsRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Type != "" {
		query.Is("type", r.Type)
	}
	if r.TeamID != "" {
		query.Is("teamId", r.TeamID)
	}
	if r.Name != "" {
		query.Is("name", r.Name)
	}
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type CreateIntegrationRequest struct {
	Type   string                 `json:"type"`
	Name   string                 `json:"name"`
	TeamID string                 `json:"teamId,omitempty"`
	Config map[string]interface{} `json:"config,omitempty"`
}

type GetIntegrationRequest struct {
	ID string
}

type UpdateIntegrationRequest struct {
	ID     string
	Name   string                 `json:"name,omitempty"`
	Config map[string]interface{} `json:"config,omitempty"`
}

type DeleteIntegrationRequest struct {
	ID string
}

