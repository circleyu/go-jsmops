package alert

import "time"

type SnoozeAlertRequest struct {
	ID        string
	Until     *time.Time `json:"until,omitempty"`
	UntilTime string     `json:"untilTime,omitempty"`
}

