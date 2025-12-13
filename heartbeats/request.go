package heartbeats

import (
	"strconv"

	"github.com/circleyu/go-jsmops/params"
)

type ListHeartbeatsRequest struct {
	TeamID string
	Name   string
	Offset int
	Size   int
}

func (r *ListHeartbeatsRequest) RequestParams() *params.Params {
	query := params.Build()
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

type CreateHeartbeatRequest struct {
	TeamID      string
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Interval    int    `json:"interval,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
}

type UpdateHeartbeatRequest struct {
	TeamID      string
	Name        string
	Description string `json:"description,omitempty"`
	Interval    int    `json:"interval,omitempty"`
	Enabled     *bool  `json:"enabled,omitempty"`
}

type DeleteHeartbeatRequest struct {
	TeamID string
	Name   string
}

type PingHeartbeatRequest struct {
	TeamID string
	Name   string
}

