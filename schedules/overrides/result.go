package overrides

type Override struct {
	Alias     string                 `json:"alias,omitempty"`
	StartTime string                 `json:"startTime,omitempty"`
	EndTime   string                 `json:"endTime,omitempty"`
	Responder map[string]interface{} `json:"responder,omitempty"`
}

type ListOverridesResult struct {
	Links     OverridesResponseLink `json:"links"`
	Overrides []Override            `json:"values"`
	Count     int                   `json:"count"`
}

type OverridesResponseLink struct {
	Next string `json:"next"`
}

