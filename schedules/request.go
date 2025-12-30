package schedules

import (
	"strconv"

	"github.com/circleyu/go-jsmops/v2/params"
)

type ListSchedulesRequest struct {
	Query  string
	Size   int
	Offset int
	Expand string
}

func (r *ListSchedulesRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Query != "" {
		query.Is("query", r.Query)
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Expand != "" {
		query.Is("expand", r.Expand)
	}
	return query
}

type CreateScheduleRequest struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	TeamID      string                 `json:"teamId,omitempty"`
	Timezone    string                 `json:"timezone,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
}

type GetScheduleRequest struct {
	ID string
}

type UpdateScheduleRequest struct {
	ID          string
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Timezone    string                 `json:"timezone,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
}

type DeleteScheduleRequest struct {
	ID string
}

