package maintenances

import (
	"strconv"

	"github.com/circleyu/go-jsmops/v2/params"
)

type ListGlobalMaintenancesRequest struct {
	Type   string
	Offset int
	Size   int
}

func (r *ListGlobalMaintenancesRequest) RequestParams() *params.Params {
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

type CreateGlobalMaintenanceRequest struct {
	Description string                 `json:"description,omitempty"`
	StartTime   string                 `json:"startTime,omitempty"`
	EndTime     string                 `json:"endTime,omitempty"`
	Rules       map[string]interface{} `json:"rules,omitempty"`
}

type GetGlobalMaintenanceRequest struct {
	ID string
}

type UpdateGlobalMaintenanceRequest struct {
	ID          string
	Description string                 `json:"description,omitempty"`
	StartTime   string                 `json:"startTime,omitempty"`
	EndTime     string                 `json:"endTime,omitempty"`
	Rules       map[string]interface{} `json:"rules,omitempty"`
}

type DeleteGlobalMaintenanceRequest struct {
	ID string
}

type CancelGlobalMaintenanceRequest struct {
	ID string
}

type ListTeamMaintenancesRequest struct {
	TeamID string
	Type   string
	Offset int
	Size   int
}

func (r *ListTeamMaintenancesRequest) RequestParams() *params.Params {
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

type CreateTeamMaintenanceRequest struct {
	TeamID      string
	Description string                 `json:"description,omitempty"`
	StartTime   string                 `json:"startTime,omitempty"`
	EndTime     string                 `json:"endTime,omitempty"`
	Rules       map[string]interface{} `json:"rules,omitempty"`
}

type GetTeamMaintenanceRequest struct {
	TeamID string
	ID     string
}

type UpdateTeamMaintenanceRequest struct {
	TeamID      string
	ID          string
	Description string                 `json:"description,omitempty"`
	StartTime   string                 `json:"startTime,omitempty"`
	EndTime     string                 `json:"endTime,omitempty"`
	Rules       map[string]interface{} `json:"rules,omitempty"`
}

type DeleteTeamMaintenanceRequest struct {
	TeamID string
	ID     string
}

type CancelTeamMaintenanceRequest struct {
	TeamID string
	ID     string
}

