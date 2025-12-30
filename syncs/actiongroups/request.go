package actiongroups

import (
	"strconv"

	"github.com/circleyu/go-jsmops/v2/params"
)

type ListSyncActionGroupsRequest struct {
	SyncID string
	Offset int
	Size   int
}

func (r *ListSyncActionGroupsRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type CreateSyncActionGroupRequest struct {
	SyncID      string
	Name        string                 `json:"name"`
	Order       int                    `json:"order,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
}

type GetSyncActionGroupRequest struct {
	SyncID string
	ID     string
}

type UpdateSyncActionGroupRequest struct {
	SyncID  string
	ID      string
	Name    string                 `json:"name,omitempty"`
	Order   int                    `json:"order,omitempty"`
	Enabled *bool                  `json:"enabled,omitempty"`
	Config  map[string]interface{} `json:"config,omitempty"`
}

type DeleteSyncActionGroupRequest struct {
	SyncID string
	ID     string
}

type ReorderSyncActionGroupRequest struct {
	SyncID string
	ID     string
	Order  int `json:"order"`
}

