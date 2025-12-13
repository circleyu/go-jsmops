package notificationrules

type NotificationRule struct {
	ID          string                 `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
}

type ListNotificationRulesResult struct {
	Links            NotificationRulesResponseLink `json:"links"`
	NotificationRules []NotificationRule          `json:"values"`
	Count             int                         `json:"count"`
}

type NotificationRulesResponseLink struct {
	Next string `json:"next"`
}

