package auditlogs

import (
	"strconv"

	"github.com/circleyu/go-jsmops/v2/params"
)

type GetAuditLogsRequest struct {
	Limit     int
	PageToken string
	Category  string
	Level     string
	StartTime string
	EndTime   string
}

func (r *GetAuditLogsRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Limit > 0 {
		query.Is("limit", strconv.Itoa(r.Limit))
	}
	if r.PageToken != "" {
		query.Is("pageToken", r.PageToken)
	}
	if r.Category != "" {
		query.Is("category", r.Category)
	}
	if r.Level != "" {
		query.Is("level", r.Level)
	}
	if r.StartTime != "" {
		query.Is("startTime", r.StartTime)
	}
	if r.EndTime != "" {
		query.Is("endTime", r.EndTime)
	}
	return query
}

