package escalations

type Escalation struct {
	ID          string            `json:"id,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Rules       []EscalationRule  `json:"rules,omitempty"`
	Enabled     bool              `json:"enabled,omitempty"`
	Repeat      *EscalationRepeat `json:"repeat,omitempty"`
}

type ListEscalationsResult struct {
	Links       EscalationsResponseLink `json:"links"`
	Escalations []Escalation            `json:"values"`
	Count       int                     `json:"count"`
}

type EscalationsResponseLink struct {
	Next string `json:"next"`
}

