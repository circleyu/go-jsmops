package roles

import (
	"strconv"

	"github.com/circleyu/go-jsmops/v2/params"
)

type ListCustomUserRolesRequest struct {
	Offset int
	Size   int
}

func (r *ListCustomUserRolesRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type GetCustomUserRoleRequest struct {
	Identifier     string
	IdentifierType string
}

func (r *GetCustomUserRoleRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.IdentifierType != "" {
		query.Is("identifierType", r.IdentifierType)
	}
	return query
}

type CreateCustomUserRoleRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

type UpdateCustomUserRoleRequest struct {
	Identifier     string
	IdentifierType string
	Name           string   `json:"name,omitempty"`
	Description    string   `json:"description,omitempty"`
	Permissions    []string `json:"permissions,omitempty"`
}

func (r *UpdateCustomUserRoleRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.IdentifierType != "" {
		query.Is("identifierType", r.IdentifierType)
	}
	return query
}

type DeleteCustomUserRoleRequest struct {
	Identifier     string
	IdentifierType string
}

func (r *DeleteCustomUserRoleRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.IdentifierType != "" {
		query.Is("identifierType", r.IdentifierType)
	}
	return query
}

type AssignCustomUserRoleRequest struct {
	RoleIdentifier string   `json:"roleIdentifier"`
	Users          []string `json:"users"`
}

