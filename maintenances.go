package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/maintenances"
)

type MaintenancesManager interface {
	ListGlobalMaintenances(*maintenances.ListGlobalMaintenancesRequest) (*maintenances.ListMaintenancesResult, error)
	CreateGlobalMaintenance(*maintenances.CreateGlobalMaintenanceRequest) (*maintenances.Maintenance, error)
	GetGlobalMaintenance(*maintenances.GetGlobalMaintenanceRequest) (*maintenances.Maintenance, error)
	UpdateGlobalMaintenance(*maintenances.UpdateGlobalMaintenanceRequest) (*maintenances.Maintenance, error)
	DeleteGlobalMaintenance(*maintenances.DeleteGlobalMaintenanceRequest) error
	CancelGlobalMaintenance(*maintenances.CancelGlobalMaintenanceRequest) (*maintenances.Maintenance, error)
	ListTeamMaintenances(*maintenances.ListTeamMaintenancesRequest) (*maintenances.ListMaintenancesResult, error)
	CreateTeamMaintenance(*maintenances.CreateTeamMaintenanceRequest) (*maintenances.Maintenance, error)
	GetTeamMaintenance(*maintenances.GetTeamMaintenanceRequest) (*maintenances.Maintenance, error)
	UpdateTeamMaintenance(*maintenances.UpdateTeamMaintenanceRequest) (*maintenances.Maintenance, error)
	DeleteTeamMaintenance(*maintenances.DeleteTeamMaintenanceRequest) error
	CancelTeamMaintenance(*maintenances.CancelTeamMaintenanceRequest) (*maintenances.Maintenance, error)
}

type maintenancesManager struct {
	*APIClient
}

func newMaintenancesManager(client *APIClient) *maintenancesManager {
	return &maintenancesManager{
		client,
	}
}

func (manager *maintenancesManager) ListGlobalMaintenances(data *maintenances.ListGlobalMaintenancesRequest) (*maintenances.ListMaintenancesResult, error) {
	output := &maintenances.ListMaintenancesResult{}
	_, err := manager.get(endpoints.maintenances.ListGlobalMaintenances, output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *maintenancesManager) CreateGlobalMaintenance(data *maintenances.CreateGlobalMaintenanceRequest) (*maintenances.Maintenance, error) {
	output := &maintenances.Maintenance{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.maintenances.CreateGlobalMaintenance, jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *maintenancesManager) GetGlobalMaintenance(data *maintenances.GetGlobalMaintenanceRequest) (*maintenances.Maintenance, error) {
	output := &maintenances.Maintenance{}
	_, err := manager.get(endpoints.maintenances.GetGlobalMaintenance(data.ID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *maintenancesManager) UpdateGlobalMaintenance(data *maintenances.UpdateGlobalMaintenanceRequest) (*maintenances.Maintenance, error) {
	output := &maintenances.Maintenance{}
	requestBody := make(map[string]interface{})
	if data.Description != "" {
		requestBody["description"] = data.Description
	}
	if data.StartTime != "" {
		requestBody["startTime"] = data.StartTime
	}
	if data.EndTime != "" {
		requestBody["endTime"] = data.EndTime
	}
	if data.Rules != nil {
		requestBody["rules"] = data.Rules
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.maintenances.UpdateGlobalMaintenance(data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *maintenancesManager) DeleteGlobalMaintenance(data *maintenances.DeleteGlobalMaintenanceRequest) error {
	return manager.delete(endpoints.maintenances.DeleteGlobalMaintenance(data.ID), nil, http.StatusNoContent, http.StatusOK)
}

func (manager *maintenancesManager) CancelGlobalMaintenance(data *maintenances.CancelGlobalMaintenanceRequest) (*maintenances.Maintenance, error) {
	output := &maintenances.Maintenance{}
	err := manager.postJSON(endpoints.maintenances.CancelGlobalMaintenance(data.ID), nil, output, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *maintenancesManager) ListTeamMaintenances(data *maintenances.ListTeamMaintenancesRequest) (*maintenances.ListMaintenancesResult, error) {
	output := &maintenances.ListMaintenancesResult{}
	_, err := manager.get(endpoints.maintenances.ListTeamMaintenances(data.TeamID), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *maintenancesManager) CreateTeamMaintenance(data *maintenances.CreateTeamMaintenanceRequest) (*maintenances.Maintenance, error) {
	output := &maintenances.Maintenance{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.maintenances.CreateTeamMaintenance(data.TeamID), jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *maintenancesManager) GetTeamMaintenance(data *maintenances.GetTeamMaintenanceRequest) (*maintenances.Maintenance, error) {
	output := &maintenances.Maintenance{}
	_, err := manager.get(endpoints.maintenances.GetTeamMaintenance(data.TeamID, data.ID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *maintenancesManager) UpdateTeamMaintenance(data *maintenances.UpdateTeamMaintenanceRequest) (*maintenances.Maintenance, error) {
	output := &maintenances.Maintenance{}
	requestBody := make(map[string]interface{})
	if data.Description != "" {
		requestBody["description"] = data.Description
	}
	if data.StartTime != "" {
		requestBody["startTime"] = data.StartTime
	}
	if data.EndTime != "" {
		requestBody["endTime"] = data.EndTime
	}
	if data.Rules != nil {
		requestBody["rules"] = data.Rules
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.maintenances.UpdateTeamMaintenance(data.TeamID, data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *maintenancesManager) DeleteTeamMaintenance(data *maintenances.DeleteTeamMaintenanceRequest) error {
	return manager.delete(endpoints.maintenances.DeleteTeamMaintenance(data.TeamID, data.ID), nil, http.StatusNoContent, http.StatusOK)
}

func (manager *maintenancesManager) CancelTeamMaintenance(data *maintenances.CancelTeamMaintenanceRequest) (*maintenances.Maintenance, error) {
	output := &maintenances.Maintenance{}
	err := manager.postJSON(endpoints.maintenances.CancelTeamMaintenance(data.TeamID, data.ID), nil, output, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

