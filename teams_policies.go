package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/teams/policies"
)

type TeamPoliciesManager interface {
	ListTeamPolicies(*policies.ListTeamPoliciesRequest) (*policies.ListTeamPoliciesResult, error)
	CreateTeamPolicy(*policies.CreateTeamPolicyRequest) (*policies.TeamPolicy, error)
	GetTeamPolicy(*policies.GetTeamPolicyRequest) (*policies.TeamPolicy, error)
	PutTeamPolicy(*policies.PutTeamPolicyRequest) (*policies.TeamPolicy, error)
	DeleteTeamPolicy(*policies.DeleteTeamPolicyRequest) error
	ChangeOrderTeamPolicy(*policies.ChangeOrderTeamPolicyRequest) (*policies.TeamPolicy, error)
	EnableTeamPolicy(*policies.EnableTeamPolicyRequest) (*policies.TeamPolicy, error)
	DisableTeamPolicy(*policies.DisableTeamPolicyRequest) (*policies.TeamPolicy, error)
}

type teamPoliciesManager struct {
	*APIClient
}

func newTeamPoliciesManager(client *APIClient) *teamPoliciesManager {
	return &teamPoliciesManager{
		client,
	}
}

func (manager *teamPoliciesManager) ListTeamPolicies(data *policies.ListTeamPoliciesRequest) (*policies.ListTeamPoliciesResult, error) {
	output := &policies.ListTeamPoliciesResult{}
	_, err := manager.get(endpoints.teamPolicies.ListTeamPolicies(data.TeamID), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *teamPoliciesManager) CreateTeamPolicy(data *policies.CreateTeamPolicyRequest) (*policies.TeamPolicy, error) {
	output := &policies.TeamPolicy{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.teamPolicies.CreateTeamPolicy(data.TeamID), jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *teamPoliciesManager) GetTeamPolicy(data *policies.GetTeamPolicyRequest) (*policies.TeamPolicy, error) {
	output := &policies.TeamPolicy{}
	_, err := manager.get(endpoints.teamPolicies.GetTeamPolicy(data.TeamID, data.PolicyID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *teamPoliciesManager) PutTeamPolicy(data *policies.PutTeamPolicyRequest) (*policies.TeamPolicy, error) {
	output := &policies.TeamPolicy{}
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
	err = manager.put(endpoints.teamPolicies.PutTeamPolicy(data.TeamID, data.PolicyID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *teamPoliciesManager) DeleteTeamPolicy(data *policies.DeleteTeamPolicyRequest) error {
	return manager.delete(endpoints.teamPolicies.DeleteTeamPolicy(data.TeamID, data.PolicyID), nil, http.StatusNoContent, http.StatusOK)
}

func (manager *teamPoliciesManager) ChangeOrderTeamPolicy(data *policies.ChangeOrderTeamPolicyRequest) (*policies.TeamPolicy, error) {
	output := &policies.TeamPolicy{}
	requestBody := map[string]interface{}{
		"order": data.Order,
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.teamPolicies.ChangeOrderTeamPolicy(data.TeamID, data.PolicyID), jsonb, output, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *teamPoliciesManager) EnableTeamPolicy(data *policies.EnableTeamPolicyRequest) (*policies.TeamPolicy, error) {
	output := &policies.TeamPolicy{}
	err := manager.postJSON(endpoints.teamPolicies.EnableTeamPolicy(data.TeamID, data.PolicyID), nil, output, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *teamPoliciesManager) DisableTeamPolicy(data *policies.DisableTeamPolicyRequest) (*policies.TeamPolicy, error) {
	output := &policies.TeamPolicy{}
	err := manager.postJSON(endpoints.teamPolicies.DisableTeamPolicy(data.TeamID, data.PolicyID), nil, output, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

