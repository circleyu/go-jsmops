package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/v2/integrations/actions"
)

type IntegrationActionsManager interface {
	ListIntegrationActions(*actions.ListIntegrationActionsRequest) (*actions.ListIntegrationActionsResult, error)
	CreateIntegrationAction(*actions.CreateIntegrationActionRequest) (*actions.IntegrationAction, error)
	GetIntegrationAction(*actions.GetIntegrationActionRequest) (*actions.IntegrationAction, error)
	UpdateIntegrationAction(*actions.UpdateIntegrationActionRequest) (*actions.IntegrationAction, error)
	DeleteIntegrationAction(*actions.DeleteIntegrationActionRequest) error
	ReorderIntegrationAction(*actions.ReorderIntegrationActionRequest) (*actions.IntegrationAction, error)
}

type integrationActionsManager struct {
	*APIClient
}

func newIntegrationActionsManager(client *APIClient) *integrationActionsManager {
	return &integrationActionsManager{
		client,
	}
}

func (manager *integrationActionsManager) ListIntegrationActions(data *actions.ListIntegrationActionsRequest) (*actions.ListIntegrationActionsResult, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &actions.ListIntegrationActionsResult{}
	_, err := manager.get(endpoints.integrationActions.ListIntegrationActions(data.IntegrationID), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *integrationActionsManager) CreateIntegrationAction(data *actions.CreateIntegrationActionRequest) (*actions.IntegrationAction, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &actions.IntegrationAction{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.integrationActions.CreateIntegrationAction(data.IntegrationID), jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *integrationActionsManager) GetIntegrationAction(data *actions.GetIntegrationActionRequest) (*actions.IntegrationAction, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &actions.IntegrationAction{}
	_, err := manager.get(endpoints.integrationActions.GetIntegrationAction(data.IntegrationID, data.ID), output, nil)
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *integrationActionsManager) UpdateIntegrationAction(data *actions.UpdateIntegrationActionRequest) (*actions.IntegrationAction, error) {
	output := &actions.IntegrationAction{}
	requestBody := make(map[string]interface{})
	if data.Type != "" {
		requestBody["type"] = data.Type
	}
	if data.Name != "" {
		requestBody["name"] = data.Name
	}
	if data.Config != nil {
		requestBody["config"] = data.Config
	}
	if data.Order > 0 {
		requestBody["order"] = data.Order
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.integrationActions.UpdateIntegrationAction(data.IntegrationID, data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *integrationActionsManager) DeleteIntegrationAction(data *actions.DeleteIntegrationActionRequest) error {
	if err := manager.checkBasicAuth(); err != nil {
		return err
	}
	return manager.delete(endpoints.integrationActions.DeleteIntegrationAction(data.IntegrationID, data.ID), nil, http.StatusNoContent, http.StatusOK)
}

func (manager *integrationActionsManager) ReorderIntegrationAction(data *actions.ReorderIntegrationActionRequest) (*actions.IntegrationAction, error) {
	output := &actions.IntegrationAction{}
	requestBody := map[string]interface{}{
		"order": data.Order,
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.integrationActions.ReorderIntegrationAction(data.IntegrationID, data.ID), jsonb, output, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

