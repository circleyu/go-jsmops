package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/v2/schedules/rotations"
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
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &rotations.ListRotationsResult{}
	_, err := manager.get(endpoints.schedulesRotations.ListRotations(data.ScheduleID), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *schedulesRotationsManager) CreateRotation(data *rotations.CreateRotationRequest) (*rotations.Rotation, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &rotations.Rotation{}
	jsonb, err := sonic.Marshal(data)
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
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &rotations.Rotation{}
	_, err := manager.get(endpoints.schedulesRotations.GetRotation(data.ScheduleID, data.ID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *schedulesRotationsManager) UpdateRotation(data *rotations.UpdateRotationRequest) (*rotations.Rotation, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
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
	jsonb, err := sonic.Marshal(requestBody)
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
	if err := manager.checkBasicAuth(); err != nil {
		return err
	}
	return manager.delete(endpoints.schedulesRotations.DeleteRotation(data.ScheduleID, data.ID), nil, http.StatusNoContent, http.StatusOK)
}

