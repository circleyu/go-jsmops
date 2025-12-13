package roles

type TeamRole struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

type ListTeamRolesResult struct {
	Links TeamRolesResponseLink `json:"links"`
	Roles []TeamRole            `json:"values"`
	Count int                   `json:"count"`
}

type TeamRolesResponseLink struct {
	Next string `json:"next"`
}

