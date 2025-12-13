package rotations

import (
	"strconv"

	"github.com/circleyu/go-jsmops/params"
)

type ListRotationsRequest struct {
	ScheduleID string
	Offset     int
	Size       int
}

func (r *ListRotationsRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type CreateRotationRequest struct {
	ScheduleID string
	Name       string                 `json:"name"`
	StartDate  string                 `json:"startDate,omitempty"`
	Length     int                    `json:"length,omitempty"`
	Participants []map[string]interface{} `json:"participants,omitempty"`
	Type       string                 `json:"type,omitempty"`
}

type GetRotationRequest struct {
	ScheduleID string
	ID         string
}

type UpdateRotationRequest struct {
	ScheduleID   string
	ID           string
	Name         string                 `json:"name,omitempty"`
	StartDate    string                 `json:"startDate,omitempty"`
	Length       int                    `json:"length,omitempty"`
	Participants []map[string]interface{} `json:"participants,omitempty"`
	Type         string                 `json:"type,omitempty"`
}

type DeleteRotationRequest struct {
	ScheduleID string
	ID         string
}

