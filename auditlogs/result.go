package auditlogs

type AuditLog struct {
	ID        string `json:"id,omitempty"`
	Category  string `json:"category,omitempty"`
	Level     string `json:"level,omitempty"`
	Message   string `json:"message,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	User      string `json:"user,omitempty"`
	Details   map[string]interface{} `json:"details,omitempty"`
}

type ListAuditLogsResult struct {
	Links AuditLogsResponseLink `json:"links"`
	Logs  []AuditLog            `json:"values"`
	Count int                   `json:"count"`
}

type AuditLogsResponseLink struct {
	Next string `json:"next"`
}

