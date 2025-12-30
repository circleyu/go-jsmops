package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/v2/syncs"
)

type SyncsManager interface {
	ListSyncs(*syncs.ListSyncsRequest) (*syncs.ListSyncsResult, error)
	CreateSync(*syncs.CreateSyncRequest) (*syncs.Sync, error)
	GetSync(*syncs.GetSyncRequest) (*syncs.Sync, error)
	UpdateSync(*syncs.UpdateSyncRequest) (*syncs.Sync, error)
	DeleteSync(*syncs.DeleteSyncRequest) error
}

type syncsManager struct {
	*APIClient
}

func newSyncsManager(client *APIClient) *syncsManager {
	return &syncsManager{
		client,
	}
}

func (manager *syncsManager) ListSyncs(data *syncs.ListSyncsRequest) (*syncs.ListSyncsResult, error) {
	output := &syncs.ListSyncsResult{}
	_, err := manager.get(endpoints.syncs.ListSyncs, output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *syncsManager) CreateSync(data *syncs.CreateSyncRequest) (*syncs.Sync, error) {
	output := &syncs.Sync{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.syncs.CreateSync, jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *syncsManager) GetSync(data *syncs.GetSyncRequest) (*syncs.Sync, error) {
	output := &syncs.Sync{}
	_, err := manager.get(endpoints.syncs.GetSync(data.ID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *syncsManager) UpdateSync(data *syncs.UpdateSyncRequest) (*syncs.Sync, error) {
	output := &syncs.Sync{}
	requestBody := make(map[string]interface{})
	if data.Name != "" {
		requestBody["name"] = data.Name
	}
	if data.Description != "" {
		requestBody["description"] = data.Description
	}
	if data.Enabled != nil {
		requestBody["enabled"] = *data.Enabled
	}
	if data.Config != nil {
		requestBody["config"] = data.Config
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.syncs.UpdateSync(data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *syncsManager) DeleteSync(data *syncs.DeleteSyncRequest) error {
	return manager.delete(endpoints.syncs.DeleteSync(data.ID), nil, http.StatusNoContent, http.StatusOK)
}

