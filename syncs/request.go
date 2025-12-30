package syncs

import (
	"strconv"

	"github.com/circleyu/go-jsmops/v2/params"
)

type ListSyncsRequest struct {
	Type   string
	TeamID string
	Offset int
	Size   int
}

func (r *ListSyncsRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Type != "" {
		query.Is("type", r.Type)
	}
	if r.TeamID != "" {
		query.Is("teamId", r.TeamID)
	}
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type CreateSyncRequest struct {
	Name        string                 `json:"name"`
	Type        string                 `json:"type"`
	TeamID      string                 `json:"teamId,omitempty"`
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
}

type GetSyncRequest struct {
	ID string
}

type UpdateSyncRequest struct {
	ID          string
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
}

type DeleteSyncRequest struct {
	ID string
}

