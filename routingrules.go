package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/v2/routingrules"
)

type RoutingRulesManager interface {
	ListRoutingRules(*routingrules.ListRoutingRulesRequest) (*routingrules.ListRoutingRulesResult, error)
	CreateRoutingRule(*routingrules.CreateRoutingRuleRequest) (*routingrules.RoutingRule, error)
	GetRoutingRule(*routingrules.GetRoutingRuleRequest) (*routingrules.RoutingRule, error)
	UpdateRoutingRule(*routingrules.UpdateRoutingRuleRequest) (*routingrules.RoutingRule, error)
	DeleteRoutingRule(*routingrules.DeleteRoutingRuleRequest) error
	ChangeOrderRoutingRule(*routingrules.ChangeOrderRoutingRuleRequest) (*routingrules.RoutingRule, error)
}

type routingRulesManager struct {
	*APIClient
}

func newRoutingRulesManager(client *APIClient) *routingRulesManager {
	return &routingRulesManager{
		client,
	}
}

func (manager *routingRulesManager) ListRoutingRules(data *routingrules.ListRoutingRulesRequest) (*routingrules.ListRoutingRulesResult, error) {
	output := &routingrules.ListRoutingRulesResult{}
	_, err := manager.get(endpoints.routingRules.ListRoutingRules(data.TeamID), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *routingRulesManager) CreateRoutingRule(data *routingrules.CreateRoutingRuleRequest) (*routingrules.RoutingRule, error) {
	output := &routingrules.RoutingRule{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.routingRules.CreateRoutingRule(data.TeamID), jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *routingRulesManager) GetRoutingRule(data *routingrules.GetRoutingRuleRequest) (*routingrules.RoutingRule, error) {
	output := &routingrules.RoutingRule{}
	_, err := manager.get(endpoints.routingRules.GetRoutingRule(data.TeamID, data.ID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *routingRulesManager) UpdateRoutingRule(data *routingrules.UpdateRoutingRuleRequest) (*routingrules.RoutingRule, error) {
	output := &routingrules.RoutingRule{}
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
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.routingRules.UpdateRoutingRule(data.TeamID, data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *routingRulesManager) DeleteRoutingRule(data *routingrules.DeleteRoutingRuleRequest) error {
	return manager.delete(endpoints.routingRules.DeleteRoutingRule(data.TeamID, data.ID), nil, http.StatusNoContent, http.StatusOK)
}

func (manager *routingRulesManager) ChangeOrderRoutingRule(data *routingrules.ChangeOrderRoutingRuleRequest) (*routingrules.RoutingRule, error) {
	output := &routingrules.RoutingRule{}
	requestBody := map[string]interface{}{
		"order": data.Order,
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.routingRules.ChangeOrderRoutingRule(data.TeamID, data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

