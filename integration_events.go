package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/v2/alert"
)

type IntegrationEventsManager interface {
	CreateAlert(*alert.IntegrationCreateAlertRequest) (*alert.SuccessResponse, error)
	AcknowledgeAlert(*alert.IntegrationAcknowledgeAlertRequest) (*alert.SuccessResponse, error)
	CloseAlert(*alert.IntegrationCloseAlertRequest) (*alert.SuccessResponse, error)
	AddNote(*alert.IntegrationAddNoteRequest) (*alert.SuccessResponse, error)
}

type integrationEventsManager struct {
	*APIClient
}

func newIntegrationEventsManager(client *APIClient) *integrationEventsManager {
	return &integrationEventsManager{
		client,
	}
}

func (manager *integrationEventsManager) CreateAlert(data *alert.IntegrationCreateAlertRequest) (*alert.SuccessResponse, error) {
	if err := manager.checkGenieKey(); err != nil {
		return nil, err
	}
	output := &alert.SuccessResponse{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postIntegrationEvent("alerts", jsonb, output, nil, http.StatusAccepted)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (manager *integrationEventsManager) AcknowledgeAlert(data *alert.IntegrationAcknowledgeAlertRequest) (*alert.SuccessResponse, error) {
	if err := manager.checkGenieKey(); err != nil {
		return nil, err
	}
	output := &alert.SuccessResponse{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	path := "alerts/" + data.IdentifierValue + "/acknowledge"
	params := data.RequestParams()
	err = manager.postIntegrationEvent(path, jsonb, output, params, http.StatusAccepted)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (manager *integrationEventsManager) CloseAlert(data *alert.IntegrationCloseAlertRequest) (*alert.SuccessResponse, error) {
	if err := manager.checkGenieKey(); err != nil {
		return nil, err
	}
	output := &alert.SuccessResponse{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	path := "alerts/" + data.IdentifierValue + "/close"
	params := data.RequestParams()
	err = manager.postIntegrationEvent(path, jsonb, output, params, http.StatusAccepted)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (manager *integrationEventsManager) AddNote(data *alert.IntegrationAddNoteRequest) (*alert.SuccessResponse, error) {
	if err := manager.checkGenieKey(); err != nil {
		return nil, err
	}
	output := &alert.SuccessResponse{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	path := "alerts/" + data.IdentifierValue + "/notes"
	params := data.RequestParams()
	err = manager.postIntegrationEvent(path, jsonb, output, params, http.StatusAccepted)
	if err != nil {
		return nil, err
	}

	return output, nil
}

