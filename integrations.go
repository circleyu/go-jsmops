package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/v2/integrations"
)

type IntegrationsManager interface {
	ListIntegrations(*integrations.ListIntegrationsRequest) (*integrations.ListIntegrationsResult, error)
	CreateIntegration(*integrations.CreateIntegrationRequest) (*integrations.Integration, error)
	GetIntegration(*integrations.GetIntegrationRequest) (*integrations.Integration, error)
	UpdateIntegration(*integrations.UpdateIntegrationRequest) (*integrations.Integration, error)
	DeleteIntegration(*integrations.DeleteIntegrationRequest) error
}

type integrationsManager struct {
	*APIClient
}

func newIntegrationsManager(client *APIClient) *integrationsManager {
	return &integrationsManager{
		client,
	}
}

func (manager *integrationsManager) ListIntegrations(data *integrations.ListIntegrationsRequest) (*integrations.ListIntegrationsResult, error) {
	output := &integrations.ListIntegrationsResult{}
	_, err := manager.get(endpoints.integrations.ListIntegrations, output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *integrationsManager) CreateIntegration(data *integrations.CreateIntegrationRequest) (*integrations.Integration, error) {
	output := &integrations.Integration{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.integrations.CreateIntegration, jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *integrationsManager) GetIntegration(data *integrations.GetIntegrationRequest) (*integrations.Integration, error) {
	output := &integrations.Integration{}
	_, err := manager.get(endpoints.integrations.GetIntegration(data.ID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *integrationsManager) UpdateIntegration(data *integrations.UpdateIntegrationRequest) (*integrations.Integration, error) {
	output := &integrations.Integration{}
	requestBody := make(map[string]interface{})
	if data.Name != "" {
		requestBody["name"] = data.Name
	}
	if data.Config != nil {
		requestBody["config"] = data.Config
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.integrations.UpdateIntegration(data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *integrationsManager) DeleteIntegration(data *integrations.DeleteIntegrationRequest) error {
	return manager.delete(endpoints.integrations.DeleteIntegration(data.ID), nil, http.StatusNoContent, http.StatusOK)
}

