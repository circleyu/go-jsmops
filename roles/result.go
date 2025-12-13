package roles

type CustomUserRole struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

type ListCustomUserRolesResult struct {
	Links RolesResponseLink `json:"links"`
	Roles []CustomUserRole  `json:"values"`
	Count int               `json:"count"`
}

type RolesResponseLink struct {
	Next string `json:"next"`
}

type SuccessResponse struct {
	Result    string  `json:"result"`
	RequestID string  `json:"requestId"`
	Took      float32 `json:"took"`
}

