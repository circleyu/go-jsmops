package timelines

import (
	"strconv"

	"github.com/circleyu/go-jsmops/params"
)

type GetScheduleTimelineRequest struct {
	ScheduleID   string
	Interval     int
	IntervalUnit string
	Date         string
	Expand       string
}

func (r *GetScheduleTimelineRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Interval > 0 {
		query.Is("interval", strconv.Itoa(r.Interval))
	}
	if r.IntervalUnit != "" {
		query.Is("intervalUnit", r.IntervalUnit)
	}
	if r.Date != "" {
		query.Is("date", r.Date)
	}
	if r.Expand != "" {
		query.Is("expand", r.Expand)
	}
	return query
}

type ExportScheduleTimelineRequest struct {
	ScheduleID string
}

