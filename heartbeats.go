package jsmops

import (
	"encoding/json"
	"net/http"

	"github.com/circleyu/go-jsmops/heartbeats"
	"github.com/circleyu/go-jsmops/params"
)

type HeartbeatsManager interface {
	ListHeartbeats(*heartbeats.ListHeartbeatsRequest) (*heartbeats.ListHeartbeatsResult, error)
	CreateHeartbeat(*heartbeats.CreateHeartbeatRequest) (*heartbeats.Heartbeat, error)
	UpdateHeartbeat(*heartbeats.UpdateHeartbeatRequest) (*heartbeats.Heartbeat, error)
	DeleteHeartbeat(*heartbeats.DeleteHeartbeatRequest) error
	PingHeartbeat(*heartbeats.PingHeartbeatRequest) (*heartbeats.PingResponse, error)
}

type heartbeatsManager struct {
	*APIClient
}

func newHeartbeatsManager(client *APIClient) *heartbeatsManager {
	return &heartbeatsManager{
		client,
	}
}

func (manager *heartbeatsManager) ListHeartbeats(data *heartbeats.ListHeartbeatsRequest) (*heartbeats.ListHeartbeatsResult, error) {
	output := &heartbeats.ListHeartbeatsResult{}
	_, err := manager.get(endpoints.heartbeats.ListHeartbeats(data.TeamID), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *heartbeatsManager) CreateHeartbeat(data *heartbeats.CreateHeartbeatRequest) (*heartbeats.Heartbeat, error) {
	output := &heartbeats.Heartbeat{}
	jsonb, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.heartbeats.CreateHeartbeat(data.TeamID), jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *heartbeatsManager) UpdateHeartbeat(data *heartbeats.UpdateHeartbeatRequest) (*heartbeats.Heartbeat, error) {
	output := &heartbeats.Heartbeat{}
	requestBody := make(map[string]interface{})
	if data.Description != "" {
		requestBody["description"] = data.Description
	}
	if data.Interval > 0 {
		requestBody["interval"] = data.Interval
	}
	if data.Enabled != nil {
		requestBody["enabled"] = *data.Enabled
	}
	jsonb, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	query := params.Build()
	query.Is("name", data.Name)
	path := endpoints.heartbeats.UpdateHeartbeat(data.TeamID) + "?" + query.URLSafe()
	err = manager.patch(path, jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *heartbeatsManager) DeleteHeartbeat(data *heartbeats.DeleteHeartbeatRequest) error {
	query := params.Build()
	query.Is("name", data.Name)
	path := endpoints.heartbeats.DeleteHeartbeat(data.TeamID) + "?" + query.URLSafe()
	return manager.delete(path, nil, http.StatusNoContent, http.StatusOK)
}

func (manager *heartbeatsManager) PingHeartbeat(data *heartbeats.PingHeartbeatRequest) (*heartbeats.PingResponse, error) {
	output := &heartbeats.PingResponse{}
	query := params.Build()
	query.Is("name", data.Name)
	path := endpoints.heartbeats.PingHeartbeat(data.TeamID) + "?" + query.URLSafe()
	err := manager.post(path, nil, output, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

