package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/syncs/actions"
)

type SyncsActionsManager interface {
	ListSyncActions(*actions.ListSyncActionsRequest) (*actions.ListSyncActionsResult, error)
	CreateSyncAction(*actions.CreateSyncActionRequest) (*actions.SyncAction, error)
	GetSyncAction(*actions.GetSyncActionRequest) (*actions.SyncAction, error)
	UpdateSyncAction(*actions.UpdateSyncActionRequest) (*actions.SyncAction, error)
	DeleteSyncAction(*actions.DeleteSyncActionRequest) error
	ReorderSyncAction(*actions.ReorderSyncActionRequest) (*actions.SyncAction, error)
}

type syncsActionsManager struct {
	*APIClient
}

func newSyncsActionsManager(client *APIClient) *syncsActionsManager {
	return &syncsActionsManager{
		client,
	}
}

func (manager *syncsActionsManager) ListSyncActions(data *actions.ListSyncActionsRequest) (*actions.ListSyncActionsResult, error) {
	output := &actions.ListSyncActionsResult{}
	_, err := manager.get(endpoints.syncsActions.ListSyncActions(data.SyncID), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *syncsActionsManager) CreateSyncAction(data *actions.CreateSyncActionRequest) (*actions.SyncAction, error) {
	output := &actions.SyncAction{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.syncsActions.CreateSyncAction(data.SyncID), jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *syncsActionsManager) GetSyncAction(data *actions.GetSyncActionRequest) (*actions.SyncAction, error) {
	output := &actions.SyncAction{}
	_, err := manager.get(endpoints.syncsActions.GetSyncAction(data.SyncID, data.ID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *syncsActionsManager) UpdateSyncAction(data *actions.UpdateSyncActionRequest) (*actions.SyncAction, error) {
	output := &actions.SyncAction{}
	requestBody := make(map[string]interface{})
	if data.Name != "" {
		requestBody["name"] = data.Name
	}
	if data.Type != "" {
		requestBody["type"] = data.Type
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
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.syncsActions.UpdateSyncAction(data.SyncID, data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *syncsActionsManager) DeleteSyncAction(data *actions.DeleteSyncActionRequest) error {
	return manager.delete(endpoints.syncsActions.DeleteSyncAction(data.SyncID, data.ID), nil, http.StatusNoContent, http.StatusOK)
}

func (manager *syncsActionsManager) ReorderSyncAction(data *actions.ReorderSyncActionRequest) (*actions.SyncAction, error) {
	output := &actions.SyncAction{}
	requestBody := map[string]interface{}{
		"order": data.Order,
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.syncsActions.ReorderSyncAction(data.SyncID, data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

