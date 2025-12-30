package jec

import (
	"strconv"

	"github.com/circleyu/go-jsmops/v2/params"
)

type ListJECChannelsRequest struct {
	Offset int
	Size   int
}

func (r *ListJECChannelsRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type CreateJECChannelRequest struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
}

type GetJECChannelRequest struct {
	ID string
}

type DeleteJECChannelRequest struct {
	ID string
}

type SendJECActionRequest struct {
	ChannelID string                 `json:"channelId,omitempty"`
	Action    string                 `json:"action"`
	Payload   map[string]interface{} `json:"payload,omitempty"`
}

