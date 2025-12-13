package steps

type NotificationRuleStep struct {
	ID          string                 `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
	Order       int                    `json:"order,omitempty"`
}

type ListNotificationRuleStepsResult struct {
	Links  NotificationRuleStepsResponseLink `json:"links"`
	Steps  []NotificationRuleStep            `json:"values"`
	Count  int                                `json:"count"`
}

type NotificationRuleStepsResponseLink struct {
	Next string `json:"next"`
}

