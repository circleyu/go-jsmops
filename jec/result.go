package jec

type JECChannel struct {
	ID          string                 `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
}

type ListJECChannelsResult struct {
	Links    JECChannelsResponseLink `json:"links"`
	Channels []JECChannel             `json:"values"`
	Count    int                      `json:"count"`
}

type JECChannelsResponseLink struct {
	Next string `json:"next"`
}

type SendJECActionResponse struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
}

