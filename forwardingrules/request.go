package forwardingrules

import (
	"strconv"

	"github.com/circleyu/go-jsmops/v2/params"
)

type ListForwardingRulesRequest struct {
	ShowAll bool
	Offset  int
	Size    int
}

func (r *ListForwardingRulesRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.ShowAll {
		query.Is("showAll", "true")
	}
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type CreateForwardingRuleRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type GetForwardingRuleRequest struct {
	ID string
}

type UpdateForwardingRuleRequest struct {
	ID   string
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type DeleteForwardingRuleRequest struct {
	ID string
}

