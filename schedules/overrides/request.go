package overrides

import (
	"strconv"

	"github.com/circleyu/go-jsmops/params"
)

type ListOverridesRequest struct {
	ScheduleID string
	Offset     int
	Size       int
}

func (r *ListOverridesRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type CreateOverrideRequest struct {
	ScheduleID string
	Alias      string                 `json:"alias,omitempty"`
	StartTime  string                 `json:"startTime,omitempty"`
	EndTime    string                 `json:"endTime,omitempty"`
	Responder  map[string]interface{} `json:"responder,omitempty"`
}

type GetOverrideRequest struct {
	ScheduleID string
	Alias      string
}

type UpdateOverrideRequest struct {
	ScheduleID string
	Alias      string
	StartTime  string                 `json:"startTime,omitempty"`
	EndTime    string                 `json:"endTime,omitempty"`
	Responder  map[string]interface{} `json:"responder,omitempty"`
}

type DeleteOverrideRequest struct {
	ScheduleID string
	Alias      string
}

