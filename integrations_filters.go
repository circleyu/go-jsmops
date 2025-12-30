package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/v2/integrations/filters"
)

type IntegrationFiltersManager interface {
	GetIntegrationAlertFilter(*filters.GetIntegrationAlertFilterRequest) (*filters.IntegrationAlertFilter, error)
	UpdateIntegrationAlertFilter(*filters.UpdateIntegrationAlertFilterRequest) (*filters.IntegrationAlertFilter, error)
}

type integrationFiltersManager struct {
	*APIClient
}

func newIntegrationFiltersManager(client *APIClient) *integrationFiltersManager {
	return &integrationFiltersManager{
		client,
	}
}

func (manager *integrationFiltersManager) GetIntegrationAlertFilter(data *filters.GetIntegrationAlertFilterRequest) (*filters.IntegrationAlertFilter, error) {
	output := &filters.IntegrationAlertFilter{}
	_, err := manager.get(endpoints.integrationFilters.GetIntegrationAlertFilter(data.IntegrationID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *integrationFiltersManager) UpdateIntegrationAlertFilter(data *filters.UpdateIntegrationAlertFilterRequest) (*filters.IntegrationAlertFilter, error) {
	output := &filters.IntegrationAlertFilter{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.put(endpoints.integrationFilters.UpdateIntegrationAlertFilter(data.IntegrationID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

