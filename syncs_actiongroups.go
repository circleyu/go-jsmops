package jsmops

import (
	"encoding/json"
	"net/http"

	"github.com/circleyu/go-jsmops/syncs/actiongroups"
)

type SyncsActionGroupsManager interface {
	ListSyncActionGroups(*actiongroups.ListSyncActionGroupsRequest) (*actiongroups.ListSyncActionGroupsResult, error)
	CreateSyncActionGroup(*actiongroups.CreateSyncActionGroupRequest) (*actiongroups.SyncActionGroup, error)
	GetSyncActionGroup(*actiongroups.GetSyncActionGroupRequest) (*actiongroups.SyncActionGroup, error)
	UpdateSyncActionGroup(*actiongroups.UpdateSyncActionGroupRequest) (*actiongroups.SyncActionGroup, error)
	DeleteSyncActionGroup(*actiongroups.DeleteSyncActionGroupRequest) error
	ReorderSyncActionGroup(*actiongroups.ReorderSyncActionGroupRequest) (*actiongroups.SyncActionGroup, error)
}

type syncsActionGroupsManager struct {
	*APIClient
}

func newSyncsActionGroupsManager(client *APIClient) *syncsActionGroupsManager {
	return &syncsActionGroupsManager{
		client,
	}
}

func (manager *syncsActionGroupsManager) ListSyncActionGroups(data *actiongroups.ListSyncActionGroupsRequest) (*actiongroups.ListSyncActionGroupsResult, error) {
	output := &actiongroups.ListSyncActionGroupsResult{}
	_, err := manager.get(endpoints.syncsActionGroups.ListSyncActionGroups(data.SyncID), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *syncsActionGroupsManager) CreateSyncActionGroup(data *actiongroups.CreateSyncActionGroupRequest) (*actiongroups.SyncActionGroup, error) {
	output := &actiongroups.SyncActionGroup{}
	jsonb, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.syncsActionGroups.CreateSyncActionGroup(data.SyncID), jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *syncsActionGroupsManager) GetSyncActionGroup(data *actiongroups.GetSyncActionGroupRequest) (*actiongroups.SyncActionGroup, error) {
	output := &actiongroups.SyncActionGroup{}
	_, err := manager.get(endpoints.syncsActionGroups.GetSyncActionGroup(data.SyncID, data.ID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *syncsActionGroupsManager) UpdateSyncActionGroup(data *actiongroups.UpdateSyncActionGroupRequest) (*actiongroups.SyncActionGroup, error) {
	output := &actiongroups.SyncActionGroup{}
	requestBody := make(map[string]interface{})
	if data.Name != "" {
		requestBody["name"] = data.Name
	}
	if data.Order > 0 {
		requestBody["order"] = data.Order
	}
	if data.Enabled != nil {
		requestBody["enabled"] = *data.Enabled
	}
	if data.Config != nil {
		requestBody["config"] = data.Config
	}
	jsonb, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.syncsActionGroups.UpdateSyncActionGroup(data.SyncID, data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *syncsActionGroupsManager) DeleteSyncActionGroup(data *actiongroups.DeleteSyncActionGroupRequest) error {
	return manager.delete(endpoints.syncsActionGroups.DeleteSyncActionGroup(data.SyncID, data.ID), nil, http.StatusNoContent, http.StatusOK)
}

func (manager *syncsActionGroupsManager) ReorderSyncActionGroup(data *actiongroups.ReorderSyncActionGroupRequest) (*actiongroups.SyncActionGroup, error) {
	output := &actiongroups.SyncActionGroup{}
	requestBody := map[string]interface{}{
		"order": data.Order,
	}
	jsonb, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.syncsActionGroups.ReorderSyncActionGroup(data.SyncID, data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

