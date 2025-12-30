package roles

import (
	"github.com/circleyu/go-jsmops/v2/params"
)

type ListTeamRolesRequest struct {
	TeamID string
}

type GetTeamRoleRequest struct {
	TeamID     string
	Identifier string
	IdentifierType string
}

func (r *GetTeamRoleRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.IdentifierType != "" {
		query.Is("identifierType", r.IdentifierType)
	}
	return query
}

type CreateTeamRoleRequest struct {
	TeamID      string
	Name        string   `json:"name"`
	Description string     `json:"description,omitempty"`
	Permissions []string   `json:"permissions,omitempty"`
}

type UpdateTeamRoleRequest struct {
	TeamID        string
	Identifier    string
	IdentifierType string
	Name          string   `json:"name,omitempty"`
	Description   string   `json:"description,omitempty"`
	Permissions   []string `json:"permissions,omitempty"`
}

func (r *UpdateTeamRoleRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.IdentifierType != "" {
		query.Is("identifierType", r.IdentifierType)
	}
	return query
}

type DeleteTeamRoleRequest struct {
	TeamID        string
	Identifier    string
	IdentifierType string
}

func (r *DeleteTeamRoleRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.IdentifierType != "" {
		query.Is("identifierType", r.IdentifierType)
	}
	return query
}

