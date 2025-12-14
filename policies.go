package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/policies"
)

type PoliciesManager interface {
	ListGlobalAlertPolicies(*policies.ListGlobalAlertPoliciesRequest) (*policies.ListPoliciesResult, error)
	CreateGlobalAlertPolicy(*policies.CreateGlobalAlertPolicyRequest) (*policies.Policy, error)
	GetGlobalAlertPolicy(*policies.GetGlobalAlertPolicyRequest) (*policies.Policy, error)
	PutGlobalAlertPolicy(*policies.PutGlobalAlertPolicyRequest) (*policies.Policy, error)
	DeleteGlobalAlertPolicy(*policies.DeleteGlobalAlertPolicyRequest) error
	ChangeOrderGlobalAlertPolicy(*policies.ChangeOrderGlobalAlertPolicyRequest) (*policies.Policy, error)
	EnableGlobalAlertPolicy(*policies.EnableGlobalAlertPolicyRequest) (*policies.Policy, error)
	DisableGlobalAlertPolicy(*policies.DisableGlobalAlertPolicyRequest) (*policies.Policy, error)
}

type policiesManager struct {
	*APIClient
}

func newPoliciesManager(client *APIClient) *policiesManager {
	return &policiesManager{
		client,
	}
}

func (manager *policiesManager) ListGlobalAlertPolicies(data *policies.ListGlobalAlertPoliciesRequest) (*policies.ListPoliciesResult, error) {
	output := &policies.ListPoliciesResult{}
	_, err := manager.get(endpoints.policies.ListGlobalAlertPolicies, output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *policiesManager) CreateGlobalAlertPolicy(data *policies.CreateGlobalAlertPolicyRequest) (*policies.Policy, error) {
	output := &policies.Policy{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.policies.CreateGlobalAlertPolicy, jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *policiesManager) GetGlobalAlertPolicy(data *policies.GetGlobalAlertPolicyRequest) (*policies.Policy, error) {
	output := &policies.Policy{}
	_, err := manager.get(endpoints.policies.GetGlobalAlertPolicy(data.PolicyID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *policiesManager) PutGlobalAlertPolicy(data *policies.PutGlobalAlertPolicyRequest) (*policies.Policy, error) {
	output := &policies.Policy{}
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
	err = manager.put(endpoints.policies.PutGlobalAlertPolicy(data.PolicyID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *policiesManager) DeleteGlobalAlertPolicy(data *policies.DeleteGlobalAlertPolicyRequest) error {
	return manager.delete(endpoints.policies.DeleteGlobalAlertPolicy(data.PolicyID), nil, http.StatusNoContent, http.StatusOK)
}

func (manager *policiesManager) ChangeOrderGlobalAlertPolicy(data *policies.ChangeOrderGlobalAlertPolicyRequest) (*policies.Policy, error) {
	output := &policies.Policy{}
	requestBody := map[string]interface{}{
		"order": data.Order,
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.policies.ChangeOrderGlobalAlertPolicy(data.PolicyID), jsonb, output, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *policiesManager) EnableGlobalAlertPolicy(data *policies.EnableGlobalAlertPolicyRequest) (*policies.Policy, error) {
	output := &policies.Policy{}
	err := manager.postJSON(endpoints.policies.EnableGlobalAlertPolicy(data.PolicyID), nil, output, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *policiesManager) DisableGlobalAlertPolicy(data *policies.DisableGlobalAlertPolicyRequest) (*policies.Policy, error) {
	output := &policies.Policy{}
	err := manager.postJSON(endpoints.policies.DisableGlobalAlertPolicy(data.PolicyID), nil, output, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

