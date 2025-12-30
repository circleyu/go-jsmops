package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/v2/escalations"
)

type EscalationsManager interface {
	ListEscalations(*escalations.ListEscalationsRequest) (*escalations.ListEscalationsResult, error)
	CreateEscalation(*escalations.CreateEscalationRequest) (*escalations.Escalation, error)
	GetEscalation(*escalations.GetEscalationRequest) (*escalations.Escalation, error)
	UpdateEscalation(*escalations.UpdateEscalationRequest) (*escalations.Escalation, error)
	DeleteEscalation(*escalations.DeleteEscalationRequest) error
}

type escalationsManager struct {
	*APIClient
}

func newEscalationsManager(client *APIClient) *escalationsManager {
	return &escalationsManager{
		client,
	}
}

func (manager *escalationsManager) ListEscalations(data *escalations.ListEscalationsRequest) (*escalations.ListEscalationsResult, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &escalations.ListEscalationsResult{}
	_, err := manager.get(endpoints.escalations.ListEscalations(data.TeamID), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *escalationsManager) CreateEscalation(data *escalations.CreateEscalationRequest) (*escalations.Escalation, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &escalations.Escalation{}
	requestBody := make(map[string]interface{})
	requestBody["name"] = data.Name
	if data.Description != "" {
		requestBody["description"] = data.Description
	}
	if len(data.Rules) > 0 {
		requestBody["rules"] = data.Rules
	}
	requestBody["enabled"] = data.Enabled
	if data.Repeat != nil {
		requestBody["repeat"] = data.Repeat
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.escalations.CreateEscalation(data.TeamID), jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *escalationsManager) GetEscalation(data *escalations.GetEscalationRequest) (*escalations.Escalation, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &escalations.Escalation{}
	_, err := manager.get(endpoints.escalations.GetEscalation(data.TeamID, data.ID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *escalationsManager) UpdateEscalation(data *escalations.UpdateEscalationRequest) (*escalations.Escalation, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &escalations.Escalation{}
	requestBody := make(map[string]interface{})
	if data.Name != "" {
		requestBody["name"] = data.Name
	}
	if data.Description != "" {
		requestBody["description"] = data.Description
	}
	if len(data.Rules) > 0 {
		requestBody["rules"] = data.Rules
	}
	if data.Enabled != nil {
		requestBody["enabled"] = *data.Enabled
	}
	if data.Repeat != nil {
		requestBody["repeat"] = data.Repeat
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.escalations.UpdateEscalation(data.TeamID, data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *escalationsManager) DeleteEscalation(data *escalations.DeleteEscalationRequest) error {
	if err := manager.checkBasicAuth(); err != nil {
		return err
	}
	return manager.delete(endpoints.escalations.DeleteEscalation(data.TeamID, data.ID), nil, http.StatusNoContent, http.StatusOK)
}

