package timelines

type ScheduleTimeline struct {
	ScheduleID string                   `json:"scheduleId,omitempty"`
	StartTime  string                   `json:"startTime,omitempty"`
	EndTime    string                   `json:"endTime,omitempty"`
	Timeline   []map[string]interface{} `json:"timeline,omitempty"`
}

