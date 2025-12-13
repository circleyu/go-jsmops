package alert

type UpdateAlertPriorityRequest struct {
	ID       string
	Priority Priority `json:"priority"`
}

