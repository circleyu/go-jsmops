package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/schedules/overrides"
)

type SchedulesOverridesManager interface {
	ListOverrides(*overrides.ListOverridesRequest) (*overrides.ListOverridesResult, error)
	CreateOverride(*overrides.CreateOverrideRequest) (*overrides.Override, error)
	GetOverride(*overrides.GetOverrideRequest) (*overrides.Override, error)
	UpdateOverride(*overrides.UpdateOverrideRequest) (*overrides.Override, error)
	DeleteOverride(*overrides.DeleteOverrideRequest) error
}

type schedulesOverridesManager struct {
	*APIClient
}

func newSchedulesOverridesManager(client *APIClient) *schedulesOverridesManager {
	return &schedulesOverridesManager{
		client,
	}
}

func (manager *schedulesOverridesManager) ListOverrides(data *overrides.ListOverridesRequest) (*overrides.ListOverridesResult, error) {
	output := &overrides.ListOverridesResult{}
	_, err := manager.get(endpoints.schedulesOverrides.ListOverrides(data.ScheduleID), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *schedulesOverridesManager) CreateOverride(data *overrides.CreateOverrideRequest) (*overrides.Override, error) {
	output := &overrides.Override{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.schedulesOverrides.CreateOverride(data.ScheduleID), jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *schedulesOverridesManager) GetOverride(data *overrides.GetOverrideRequest) (*overrides.Override, error) {
	output := &overrides.Override{}
	_, err := manager.get(endpoints.schedulesOverrides.GetOverride(data.ScheduleID, data.Alias), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *schedulesOverridesManager) UpdateOverride(data *overrides.UpdateOverrideRequest) (*overrides.Override, error) {
	output := &overrides.Override{}
	requestBody := make(map[string]interface{})
	if data.StartTime != "" {
		requestBody["startTime"] = data.StartTime
	}
	if data.EndTime != "" {
		requestBody["endTime"] = data.EndTime
	}
	if data.Responder != nil {
		requestBody["responder"] = data.Responder
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.put(endpoints.schedulesOverrides.UpdateOverride(data.ScheduleID, data.Alias), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *schedulesOverridesManager) DeleteOverride(data *overrides.DeleteOverrideRequest) error {
	return manager.delete(endpoints.schedulesOverrides.DeleteOverride(data.ScheduleID, data.Alias), nil, http.StatusNoContent, http.StatusOK)
}

