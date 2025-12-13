package alert

type EscalateAlertRequest struct {
	ID      string
	Escalation EscalationRequest `json:"escalation,omitempty"`
}

type EscalationRequest struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

