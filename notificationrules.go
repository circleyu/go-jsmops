package jsmops

import (
	"encoding/json"
	"net/http"

	"github.com/circleyu/go-jsmops/notificationrules"
)

type NotificationRulesManager interface {
	ListNotificationRules(*notificationrules.ListNotificationRulesRequest) (*notificationrules.ListNotificationRulesResult, error)
	CreateNotificationRule(*notificationrules.CreateNotificationRuleRequest) (*notificationrules.NotificationRule, error)
	GetNotificationRule(*notificationrules.GetNotificationRuleRequest) (*notificationrules.NotificationRule, error)
	UpdateNotificationRule(*notificationrules.UpdateNotificationRuleRequest) (*notificationrules.NotificationRule, error)
	DeleteNotificationRule(*notificationrules.DeleteNotificationRuleRequest) error
}

type notificationRulesManager struct {
	*APIClient
}

func newNotificationRulesManager(client *APIClient) *notificationRulesManager {
	return &notificationRulesManager{
		client,
	}
}

func (manager *notificationRulesManager) ListNotificationRules(data *notificationrules.ListNotificationRulesRequest) (*notificationrules.ListNotificationRulesResult, error) {
	output := &notificationrules.ListNotificationRulesResult{}
	_, err := manager.get(endpoints.notificationRules.ListNotificationRules, output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *notificationRulesManager) CreateNotificationRule(data *notificationrules.CreateNotificationRuleRequest) (*notificationrules.NotificationRule, error) {
	output := &notificationrules.NotificationRule{}
	jsonb, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.notificationRules.CreateNotificationRule, jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *notificationRulesManager) GetNotificationRule(data *notificationrules.GetNotificationRuleRequest) (*notificationrules.NotificationRule, error) {
	output := &notificationrules.NotificationRule{}
	_, err := manager.get(endpoints.notificationRules.GetNotificationRule(data.ID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *notificationRulesManager) UpdateNotificationRule(data *notificationrules.UpdateNotificationRuleRequest) (*notificationrules.NotificationRule, error) {
	output := &notificationrules.NotificationRule{}
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
	jsonb, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.notificationRules.UpdateNotificationRule(data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *notificationRulesManager) DeleteNotificationRule(data *notificationrules.DeleteNotificationRuleRequest) error {
	return manager.delete(endpoints.notificationRules.DeleteNotificationRule(data.ID), nil, http.StatusNoContent, http.StatusOK)
}

