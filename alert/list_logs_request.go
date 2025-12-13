package alert

import (
	"strconv"

	"github.com/circleyu/go-jsmops/params"
)

type ListAlertLogsRequest struct {
	ID    string
	After string
	Size  int
}

func (r *ListAlertLogsRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.After != "" {
		query.Is("after", r.After)
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

