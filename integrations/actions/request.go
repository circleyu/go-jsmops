package actions

import (
	"strconv"

	"github.com/circleyu/go-jsmops/v2/params"
)

type ListIntegrationActionsRequest struct {
	IntegrationID string
	Offset        int
	Size          int
}

func (r *ListIntegrationActionsRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type CreateIntegrationActionRequest struct {
	IntegrationID string
	Type           string                 `json:"type"`
	Name           string                 `json:"name"`
	Config         map[string]interface{} `json:"config,omitempty"`
	Order          int                    `json:"order,omitempty"`
}

type GetIntegrationActionRequest struct {
	IntegrationID string
	ID             string
}

type UpdateIntegrationActionRequest struct {
	IntegrationID string
	ID             string
	Type           string                 `json:"type,omitempty"`
	Name           string                 `json:"name,omitempty"`
	Config         map[string]interface{} `json:"config,omitempty"`
	Order          int                    `json:"order,omitempty"`
}

type DeleteIntegrationActionRequest struct {
	IntegrationID string
	ID             string
}

type ReorderIntegrationActionRequest struct {
	IntegrationID string
	ID             string
	Order          int `json:"order"`
}

