package alert

import (
	"strconv"

	"github.com/circleyu/go-jsmops/v2/params"
)

type ListAlertsRequest struct {
	Limit                int
	Sort                 SortField
	Offset               int
	Order                Order
	Query                string
	SearchIdentifier     string
	SearchIdentifierType SearchIdentifierType
}

func (r *ListAlertsRequest) RequestParams() *params.Params {
	query := params.Build()

	if r.Limit != 0 {
		query.Is("limit", strconv.Itoa(r.Limit))
	}

	if r.Sort != "" {
		query.Is("sort", string(r.Sort))
	}

	if r.Offset != 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}

	if r.Query != "" {
		query.Is("query", r.Query)
	}

	if r.SearchIdentifier != "" {
		query.Is("searchIdentifier", r.SearchIdentifier)
	}

	if r.SearchIdentifierType != "" {
		query.Is("searchIdentifierType", string(r.SearchIdentifierType))
	}

	if r.Order != "" {
		query.Is("order", string(r.Order))
	}

	return query
}
