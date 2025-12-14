package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/roles"
)

type RolesManager interface {
	ListCustomUserRoles(*roles.ListCustomUserRolesRequest) (*roles.ListCustomUserRolesResult, error)
	GetCustomUserRole(*roles.GetCustomUserRoleRequest) (*roles.CustomUserRole, error)
	CreateCustomUserRole(*roles.CreateCustomUserRoleRequest) (*roles.CustomUserRole, error)
	UpdateCustomUserRole(*roles.UpdateCustomUserRoleRequest) (*roles.CustomUserRole, error)
	DeleteCustomUserRole(*roles.DeleteCustomUserRoleRequest) error
	AssignCustomUserRole(*roles.AssignCustomUserRoleRequest) (*roles.SuccessResponse, error)
}

type rolesManager struct {
	*APIClient
}

func newRolesManager(client *APIClient) *rolesManager {
	return &rolesManager{
		client,
	}
}

func (manager *rolesManager) ListCustomUserRoles(data *roles.ListCustomUserRolesRequest) (*roles.ListCustomUserRolesResult, error) {
	output := &roles.ListCustomUserRolesResult{}
	_, err := manager.get(endpoints.roles.ListCustomUserRoles, output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *rolesManager) GetCustomUserRole(data *roles.GetCustomUserRoleRequest) (*roles.CustomUserRole, error) {
	output := &roles.CustomUserRole{}
	_, err := manager.get(endpoints.roles.GetCustomUserRole(data.Identifier), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *rolesManager) CreateCustomUserRole(data *roles.CreateCustomUserRoleRequest) (*roles.CustomUserRole, error) {
	output := &roles.CustomUserRole{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.roles.CreateCustomUserRole, jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *rolesManager) UpdateCustomUserRole(data *roles.UpdateCustomUserRoleRequest) (*roles.CustomUserRole, error) {
	output := &roles.CustomUserRole{}
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
		path = endpoints.roles.UpdateCustomUserRole(data.Identifier) + "?" + params.URLSafe()
	} else {
		path = endpoints.roles.UpdateCustomUserRole(data.Identifier)
	}
	err = manager.put(path, jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *rolesManager) DeleteCustomUserRole(data *roles.DeleteCustomUserRoleRequest) error {
	params := data.RequestParams()
	var path string
	if params != nil {
		path = endpoints.roles.DeleteCustomUserRole(data.Identifier) + "?" + params.URLSafe()
	} else {
		path = endpoints.roles.DeleteCustomUserRole(data.Identifier)
	}
	return manager.delete(path, nil, http.StatusNoContent, http.StatusOK)
}

func (manager *rolesManager) AssignCustomUserRole(data *roles.AssignCustomUserRoleRequest) (*roles.SuccessResponse, error) {
	output := &roles.SuccessResponse{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.roles.AssignCustomUserRole, jsonb, output, nil, http.StatusOK, http.StatusAccepted)
	if err != nil {
		return nil, err
	}
	return output, nil
}

