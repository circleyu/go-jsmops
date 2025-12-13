package teams

type ListTeamsRequest struct{}

type EnableOpsRequest struct {
	TeamID string
}

type GetTeamRequestStatusRequest struct {
	TeamID    string
	RequestID string
}

