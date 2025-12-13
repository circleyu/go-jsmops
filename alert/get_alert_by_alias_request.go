package alert

import (
	"github.com/circleyu/go-jsmops/params"
)

type GetAlertByAliasRequest struct {
	Alias string
}

func (r *GetAlertByAliasRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Alias != "" {
		query.Is("alias", r.Alias)
	}
	return query
}

