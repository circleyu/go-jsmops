package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/v2/teams/roles"
)

type TeamRolesManager interface {
	ListTeamRoles(*roles.ListTeamRolesRequest) (*roles.ListTeamRolesResult, error)
	GetTeamRole(*roles.GetTeamRoleRequest) (*roles.TeamRole, error)
	CreateTeamRole(*roles.CreateTeamRoleRequest) (*roles.TeamRole, error)
	UpdateTeamRole(*roles.UpdateTeamRoleRequest) (*roles.TeamRole, error)
	DeleteTeamRole(*roles.DeleteTeamRoleRequest) error
}

type teamRolesManager struct {
	*APIClient
}

func newTeamRolesManager(client *APIClient) *teamRolesManager {
	return &teamRolesManager{
		client,
	}
}

func (manager *teamRolesManager) ListTeamRoles(data *roles.ListTeamRolesRequest) (*roles.ListTeamRolesResult, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &roles.ListTeamRolesResult{}
	_, err := manager.get(endpoints.teamRoles.ListTeamRoles(data.TeamID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *teamRolesManager) GetTeamRole(data *roles.GetTeamRoleRequest) (*roles.TeamRole, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &roles.TeamRole{}
	_, err := manager.get(endpoints.teamRoles.GetTeamRole(data.TeamID, data.Identifier), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *teamRolesManager) CreateTeamRole(data *roles.CreateTeamRoleRequest) (*roles.TeamRole, error) {
	output := &roles.TeamRole{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.teamRoles.CreateTeamRole(data.TeamID), jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *teamRolesManager) UpdateTeamRole(data *roles.UpdateTeamRoleRequest) (*roles.TeamRole, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &roles.TeamRole{}
	requestBody := make(map[string]interface{})
	if data.Name != "" {
		requestBody["name"] = data.Name
	}
	if data.Description != "" {
		requestBody["description"] = data.Description
	}
	if len(data.Permissions) > 0 {
		requestBody["permissions"] = data.Permissions
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	params := data.RequestParams()
	var path string
	if params != nil {
		path = endpoints.teamRoles.UpdateTeamRole(data.TeamID, data.Identifier) + "?" + params.URLSafe()
	} else {
		path = endpoints.teamRoles.UpdateTeamRole(data.TeamID, data.Identifier)
	}
	err = manager.patch(path, jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *teamRolesManager) DeleteTeamRole(data *roles.DeleteTeamRoleRequest) error {
	if err := manager.checkBasicAuth(); err != nil {
		return err
	}
	params := data.RequestParams()
	var path string
	if params != nil {
		path = endpoints.teamRoles.DeleteTeamRole(data.TeamID, data.Identifier) + "?" + params.URLSafe()
	} else {
		path = endpoints.teamRoles.DeleteTeamRole(data.TeamID, data.Identifier)
	}
	return manager.delete(path, nil, http.StatusNoContent, http.StatusOK)
}

