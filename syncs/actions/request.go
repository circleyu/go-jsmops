package actions

import (
	"strconv"

	"github.com/circleyu/go-jsmops/v2/params"
)

type ListSyncActionsRequest struct {
	SyncID string
	Offset int
	Size   int
}

func (r *ListSyncActionsRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type CreateSyncActionRequest struct {
	SyncID      string
	Name        string                 `json:"name"`
	Type        string                 `json:"type"`
	Order       int                    `json:"order,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
}

type GetSyncActionRequest struct {
	SyncID string
	ID     string
}

type UpdateSyncActionRequest struct {
	SyncID  string
	ID      string
	Name    string                 `json:"name,omitempty"`
	Type    string                 `json:"type,omitempty"`
	Order   int                    `json:"order,omitempty"`
	Enabled *bool                  `json:"enabled,omitempty"`
	Config  map[string]interface{} `json:"config,omitempty"`
}

type DeleteSyncActionRequest struct {
	SyncID string
	ID     string
}

type ReorderSyncActionRequest struct {
	SyncID string
	ID     string
	Order  int `json:"order"`
}

