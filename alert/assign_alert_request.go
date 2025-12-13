package alert

type AssignAlertRequest struct {
	ID        string
	Owner     *Responder `json:"owner,omitempty"`
	OwnerTeam string     `json:"ownerTeam,omitempty"`
}

