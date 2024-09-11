package jsmops

import (
	"encoding/json"
	"net/http"

	"github.com/circleyu/go-jsmops/alert"
	"github.com/circleyu/go-jsmops/querybuilder"
)

type AlertsManager interface {
	CreateAlert(*alert.CreateAlertRequest) (*alert.SuccessResponse, error)
	ListAlerts(*alert.ListAlertsRequest, *querybuilder.Query) (*alert.ListAlertsResult, error)
	AcknowledgeAlert(*alert.AcknowledgeAlertRequest) (*alert.SuccessResponse, error)
	CloseAlert(*alert.CloseAlertRequest) (*alert.SuccessResponse, error)
	AddAlertNote(*alert.AddNoteRequest) (*alert.SuccessResponse, error)
}

type alertsManager struct {
	*APIClient
}

func newAlertsManager(client *APIClient) *alertsManager {
	return &alertsManager{
		client,
	}
}

func (manager *alertsManager) CreateAlert(data *alert.CreateAlertRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	jsonb, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.alerts.CreateAlert, jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (manager *alertsManager) AcknowledgeAlert(data *alert.AcknowledgeAlertRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	jsonb, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.alerts.AcknowledgeAlert(data.IdentifierValue), jsonb, output, http.StatusAccepted)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (manager *alertsManager) CloseAlert(data *alert.CloseAlertRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	jsonb, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.alerts.CloseAlert(data.IdentifierValue), jsonb, output, http.StatusAccepted)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (manager *alertsManager) AddAlertNote(data *alert.AddNoteRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	jsonb, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.alerts.CreateAlert, jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (manager *alertsManager) ListAlerts(data *alert.ListAlertsRequest, query *querybuilder.Query) (*alert.ListAlertsResult, error) {
	output := &alert.ListAlertsResult{}

	_, err := manager.get(endpoints.alerts.ListAlerts, output, data.RequestParams())
	if err != nil {
		return nil, err
	}

	return output, nil
}
