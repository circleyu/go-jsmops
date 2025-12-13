package oncalls

type OnCallResponder struct {
	UserID   string `json:"userId,omitempty"`
	Username string `json:"username,omitempty"`
	Start    string `json:"start,omitempty"`
	End      string `json:"end,omitempty"`
}

type ListOnCallRespondersResult struct {
	Links     OnCallRespondersResponseLink `json:"links"`
	Responders []OnCallResponder            `json:"values"`
	Count      int                          `json:"count"`
}

type OnCallRespondersResponseLink struct {
	Next string `json:"next"`
}

