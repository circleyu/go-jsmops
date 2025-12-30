package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/v2/forwardingrules"
)

type ForwardingRulesManager interface {
	ListForwardingRules(*forwardingrules.ListForwardingRulesRequest) (*forwardingrules.ListForwardingRulesResult, error)
	CreateForwardingRule(*forwardingrules.CreateForwardingRuleRequest) (*forwardingrules.ForwardingRule, error)
	GetForwardingRule(*forwardingrules.GetForwardingRuleRequest) (*forwardingrules.ForwardingRule, error)
	UpdateForwardingRule(*forwardingrules.UpdateForwardingRuleRequest) (*forwardingrules.ForwardingRule, error)
	DeleteForwardingRule(*forwardingrules.DeleteForwardingRuleRequest) error
}

type forwardingRulesManager struct {
	*APIClient
}

func newForwardingRulesManager(client *APIClient) *forwardingRulesManager {
	return &forwardingRulesManager{
		client,
	}
}

func (manager *forwardingRulesManager) ListForwardingRules(data *forwardingrules.ListForwardingRulesRequest) (*forwardingrules.ListForwardingRulesResult, error) {
	output := &forwardingrules.ListForwardingRulesResult{}
	_, err := manager.get(endpoints.forwardingRules.ListForwardingRules, output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *forwardingRulesManager) CreateForwardingRule(data *forwardingrules.CreateForwardingRuleRequest) (*forwardingrules.ForwardingRule, error) {
	output := &forwardingrules.ForwardingRule{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.forwardingRules.CreateForwardingRule, jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *forwardingRulesManager) GetForwardingRule(data *forwardingrules.GetForwardingRuleRequest) (*forwardingrules.ForwardingRule, error) {
	output := &forwardingrules.ForwardingRule{}
	_, err := manager.get(endpoints.forwardingRules.GetForwardingRule(data.ID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *forwardingRulesManager) UpdateForwardingRule(data *forwardingrules.UpdateForwardingRuleRequest) (*forwardingrules.ForwardingRule, error) {
	output := &forwardingrules.ForwardingRule{}
	requestBody := make(map[string]interface{})
	if data.From != "" {
		requestBody["from"] = data.From
	}
	if data.To != "" {
		requestBody["to"] = data.To
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.put(endpoints.forwardingRules.UpdateForwardingRule(data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *forwardingRulesManager) DeleteForwardingRule(data *forwardingrules.DeleteForwardingRuleRequest) error {
	return manager.delete(endpoints.forwardingRules.DeleteForwardingRule(data.ID), nil, http.StatusNoContent, http.StatusOK)
}

