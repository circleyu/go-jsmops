package jsmops

import (
	"encoding/json"
	"net/http"

	"github.com/circleyu/go-jsmops/schedules/rotations"
)

type SchedulesRotationsManager interface {
	ListRotations(*rotations.ListRotationsRequest) (*rotations.ListRotationsResult, error)
	CreateRotation(*rotations.CreateRotationRequest) (*rotations.Rotation, error)
	GetRotation(*rotations.GetRotationRequest) (*rotations.Rotation, error)
	UpdateRotation(*rotations.UpdateRotationRequest) (*rotations.Rotation, error)
	DeleteRotation(*rotations.DeleteRotationRequest) error
}

type schedulesRotationsManager struct {
	*APIClient
}

func newSchedulesRotationsManager(client *APIClient) *schedulesRotationsManager {
	return &schedulesRotationsManager{
		client,
	}
}

func (manager *schedulesRotationsManager) ListRotations(data *rotations.ListRotationsRequest) (*rotations.ListRotationsResult, error) {
	output := &rotations.ListRotationsResult{}
	_, err := manager.get(endpoints.schedulesRotations.ListRotations(data.ScheduleID), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *schedulesRotationsManager) CreateRotation(data *rotations.CreateRotationRequest) (*rotations.Rotation, error) {
	output := &rotations.Rotation{}
	jsonb, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.schedulesRotations.CreateRotation(data.ScheduleID), jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *schedulesRotationsManager) GetRotation(data *rotations.GetRotationRequest) (*rotations.Rotation, error) {
	output := &rotations.Rotation{}
	_, err := manager.get(endpoints.schedulesRotations.GetRotation(data.ScheduleID, data.ID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *schedulesRotationsManager) UpdateRotation(data *rotations.UpdateRotationRequest) (*rotations.Rotation, error) {
	output := &rotations.Rotation{}
	requestBody := make(map[string]interface{})
	if data.Name != "" {
		requestBody["name"] = data.Name
	}
	if data.StartDate != "" {
		requestBody["startDate"] = data.StartDate
	}
	if data.Length > 0 {
		requestBody["length"] = data.Length
	}
	if len(data.Participants) > 0 {
		requestBody["participants"] = data.Participants
	}
	if data.Type != "" {
		requestBody["type"] = data.Type
	}
	jsonb, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.schedulesRotations.UpdateRotation(data.ScheduleID, data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *schedulesRotationsManager) DeleteRotation(data *rotations.DeleteRotationRequest) error {
	return manager.delete(endpoints.schedulesRotations.DeleteRotation(data.ScheduleID, data.ID), nil, http.StatusNoContent, http.StatusOK)
}

