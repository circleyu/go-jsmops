package alert

import (
	"strconv"

	"github.com/circleyu/go-jsmops/v2/params"
)

type ListAlertNotesRequest struct {
	ID     string
	After  string
	Size   int
}

func (r *ListAlertNotesRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.After != "" {
		query.Is("after", r.After)
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

