package alert

type UpdateAlertMessageRequest struct {
	ID      string
	Message string `json:"message"`
}

