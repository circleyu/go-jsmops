package alert

type UpdateAlertDescriptionRequest struct {
	ID          string
	Description string `json:"description"`
}

