package teams

type Team struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ListTeamsResult struct {
	Teams []Team `json:"values"`
	Count int    `json:"count"`
}

type SuccessResponse struct {
	Result    string  `json:"result"`
	RequestID string  `json:"requestId"`
	Took      float32 `json:"took"`
}

type RequestStatusResponse struct {
	Status    string `json:"status,omitempty"`
	RequestID string `json:"requestId,omitempty"`
}

