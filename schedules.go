package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/v2/schedules"
)

type SchedulesManager interface {
	ListSchedules(*schedules.ListSchedulesRequest) (*schedules.ListSchedulesResult, error)
	CreateSchedule(*schedules.CreateScheduleRequest) (*schedules.Schedule, error)
	GetSchedule(*schedules.GetScheduleRequest) (*schedules.Schedule, error)
	UpdateSchedule(*schedules.UpdateScheduleRequest) (*schedules.Schedule, error)
	DeleteSchedule(*schedules.DeleteScheduleRequest) error
}

type schedulesManager struct {
	*APIClient
}

func newSchedulesManager(client *APIClient) *schedulesManager {
	return &schedulesManager{
		client,
	}
}

func (manager *schedulesManager) ListSchedules(data *schedules.ListSchedulesRequest) (*schedules.ListSchedulesResult, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &schedules.ListSchedulesResult{}
	_, err := manager.get(endpoints.schedules.ListSchedules, output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *schedulesManager) CreateSchedule(data *schedules.CreateScheduleRequest) (*schedules.Schedule, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &schedules.Schedule{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.schedules.CreateSchedule, jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *schedulesManager) GetSchedule(data *schedules.GetScheduleRequest) (*schedules.Schedule, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &schedules.Schedule{}
	_, err := manager.get(endpoints.schedules.GetSchedule(data.ID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *schedulesManager) UpdateSchedule(data *schedules.UpdateScheduleRequest) (*schedules.Schedule, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &schedules.Schedule{}
	requestBody := make(map[string]interface{})
	if data.Name != "" {
		requestBody["name"] = data.Name
	}
	if data.Description != "" {
		requestBody["description"] = data.Description
	}
	if data.Timezone != "" {
		requestBody["timezone"] = data.Timezone
	}
	if data.Enabled != nil {
		requestBody["enabled"] = *data.Enabled
	}
	if data.Config != nil {
		requestBody["config"] = data.Config
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.schedules.UpdateSchedule(data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *schedulesManager) DeleteSchedule(data *schedules.DeleteScheduleRequest) error {
	if err := manager.checkBasicAuth(); err != nil {
		return err
	}
	return manager.delete(endpoints.schedules.DeleteSchedule(data.ID), nil, http.StatusNoContent, http.StatusOK)
}

