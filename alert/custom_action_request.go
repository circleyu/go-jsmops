package alert

type ExecuteCustomActionRequest struct {
	ID     string
	Action string            `json:"action"`
	User   string            `json:"user,omitempty"`
	Note   string            `json:"note,omitempty"`
}

