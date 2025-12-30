package jsmops

import (
	"bytes"
	"fmt"

	"github.com/circleyu/go-jsmops/v2/schedules/timelines"
)

type SchedulesTimelinesManager interface {
	GetScheduleTimeline(*timelines.GetScheduleTimelineRequest) (*timelines.ScheduleTimeline, error)
	ExportScheduleTimeline(*timelines.ExportScheduleTimelineRequest) (*bytes.Reader, error)
}

type schedulesTimelinesManager struct {
	*APIClient
}

func newSchedulesTimelinesManager(client *APIClient) *schedulesTimelinesManager {
	return &schedulesTimelinesManager{
		client,
	}
}

func (manager *schedulesTimelinesManager) GetScheduleTimeline(data *timelines.GetScheduleTimelineRequest) (*timelines.ScheduleTimeline, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &timelines.ScheduleTimeline{}
	_, err := manager.get(endpoints.schedulesTimelines.GetScheduleTimeline(data.ScheduleID), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *schedulesTimelinesManager) ExportScheduleTimeline(data *timelines.ExportScheduleTimelineRequest) (*bytes.Reader, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("https://api.atlassian.com/jsm/ops/api/%s/%s", manager.cloudID, endpoints.schedulesTimelines.ExportScheduleTimeline(data.ScheduleID))
	return manager.getFile(path)
}

