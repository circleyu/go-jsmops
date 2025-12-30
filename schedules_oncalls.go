package jsmops

import (
	"bytes"
	"fmt"

	"github.com/circleyu/go-jsmops/v2/schedules/oncalls"
)

type SchedulesOnCallsManager interface {
	ListOnCallResponders(*oncalls.ListOnCallRespondersRequest) (*oncalls.ListOnCallRespondersResult, error)
	ListNextOnCallResponders(*oncalls.ListNextOnCallRespondersRequest) (*oncalls.ListOnCallRespondersResult, error)
	ExportOnCallResponders(*oncalls.ExportOnCallRespondersRequest) (*bytes.Reader, error)
}

type schedulesOnCallsManager struct {
	*APIClient
}

func newSchedulesOnCallsManager(client *APIClient) *schedulesOnCallsManager {
	return &schedulesOnCallsManager{
		client,
	}
}

func (manager *schedulesOnCallsManager) ListOnCallResponders(data *oncalls.ListOnCallRespondersRequest) (*oncalls.ListOnCallRespondersResult, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &oncalls.ListOnCallRespondersResult{}
	_, err := manager.get(endpoints.schedulesOnCalls.ListOnCallResponders(data.ScheduleID), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *schedulesOnCallsManager) ListNextOnCallResponders(data *oncalls.ListNextOnCallRespondersRequest) (*oncalls.ListOnCallRespondersResult, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &oncalls.ListOnCallRespondersResult{}
	_, err := manager.get(endpoints.schedulesOnCalls.ListNextOnCallResponders(data.ScheduleID), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *schedulesOnCallsManager) ExportOnCallResponders(data *oncalls.ExportOnCallRespondersRequest) (*bytes.Reader, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("https://api.atlassian.com/jsm/ops/api/%s/%s", manager.cloudID, endpoints.schedulesOnCalls.ExportOnCallResponders(data.UserIdentifier))
	return manager.getFile(path)
}

