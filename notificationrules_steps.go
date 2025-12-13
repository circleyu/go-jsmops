package jsmops

import (
	"encoding/json"
	"net/http"

	"github.com/circleyu/go-jsmops/notificationrules/steps"
)

type NotificationRuleStepsManager interface {
	ListNotificationRuleSteps(*steps.ListNotificationRuleStepsRequest) (*steps.ListNotificationRuleStepsResult, error)
	CreateNotificationRuleStep(*steps.CreateNotificationRuleStepRequest) (*steps.NotificationRuleStep, error)
	GetNotificationRuleStep(*steps.GetNotificationRuleStepRequest) (*steps.NotificationRuleStep, error)
	UpdateNotificationRuleStep(*steps.UpdateNotificationRuleStepRequest) (*steps.NotificationRuleStep, error)
	DeleteNotificationRuleStep(*steps.DeleteNotificationRuleStepRequest) error
}

type notificationRuleStepsManager struct {
	*APIClient
}

func newNotificationRuleStepsManager(client *APIClient) *notificationRuleStepsManager {
	return &notificationRuleStepsManager{
		client,
	}
}

func (manager *notificationRuleStepsManager) ListNotificationRuleSteps(data *steps.ListNotificationRuleStepsRequest) (*steps.ListNotificationRuleStepsResult, error) {
	output := &steps.ListNotificationRuleStepsResult{}
	_, err := manager.get(endpoints.notificationRuleSteps.ListNotificationRuleSteps(data.RuleID), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *notificationRuleStepsManager) CreateNotificationRuleStep(data *steps.CreateNotificationRuleStepRequest) (*steps.NotificationRuleStep, error) {
	output := &steps.NotificationRuleStep{}
	jsonb, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.notificationRuleSteps.CreateNotificationRuleStep(data.RuleID), jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *notificationRuleStepsManager) GetNotificationRuleStep(data *steps.GetNotificationRuleStepRequest) (*steps.NotificationRuleStep, error) {
	output := &steps.NotificationRuleStep{}
	_, err := manager.get(endpoints.notificationRuleSteps.GetNotificationRuleStep(data.RuleID, data.ID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *notificationRuleStepsManager) UpdateNotificationRuleStep(data *steps.UpdateNotificationRuleStepRequest) (*steps.NotificationRuleStep, error) {
	output := &steps.NotificationRuleStep{}
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
	if data.Order > 0 {
		requestBody["order"] = data.Order
	}
	jsonb, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.notificationRuleSteps.UpdateNotificationRuleStep(data.RuleID, data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *notificationRuleStepsManager) DeleteNotificationRuleStep(data *steps.DeleteNotificationRuleStepRequest) error {
	return manager.delete(endpoints.notificationRuleSteps.DeleteNotificationRuleStep(data.RuleID, data.ID), nil, http.StatusNoContent, http.StatusOK)
}

