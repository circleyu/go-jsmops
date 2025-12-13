package oncalls

import (
	"strconv"

	"github.com/circleyu/go-jsmops/params"
)

type ListOnCallRespondersRequest struct {
	ScheduleID string
	Offset     int
	Size       int
}

func (r *ListOnCallRespondersRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type ListNextOnCallRespondersRequest struct {
	ScheduleID string
	Offset     int
	Size       int
}

func (r *ListNextOnCallRespondersRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type ExportOnCallRespondersRequest struct {
	UserIdentifier string
}

