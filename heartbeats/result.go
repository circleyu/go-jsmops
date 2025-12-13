package heartbeats

type Heartbeat struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Interval    int    `json:"interval,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
}

type ListHeartbeatsResult struct {
	Links      HeartbeatsResponseLink `json:"links"`
	Heartbeats []Heartbeat            `json:"values"`
	Count      int                    `json:"count"`
}

type HeartbeatsResponseLink struct {
	Next string `json:"next"`
}

type SuccessResponse struct {
	Result    string  `json:"result"`
	RequestID string  `json:"requestId"`
	Took      float32 `json:"took"`
}

type PingResponse struct {
	Result string `json:"result"`
}

