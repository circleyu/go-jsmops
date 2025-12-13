package jsmops

import (
	"net/http"

	"github.com/circleyu/go-jsmops/teams"
)

type TeamsManager interface {
	ListTeams(*teams.ListTeamsRequest) (*teams.ListTeamsResult, error)
	EnableOps(*teams.EnableOpsRequest) (*teams.SuccessResponse, error)
	GetTeamRequestStatus(*teams.GetTeamRequestStatusRequest) (*teams.RequestStatusResponse, error)
}

type teamsManager struct {
	*APIClient
}

func newTeamsManager(client *APIClient) *teamsManager {
	return &teamsManager{
		client,
	}
}

func (manager *teamsManager) ListTeams(data *teams.ListTeamsRequest) (*teams.ListTeamsResult, error) {
	output := &teams.ListTeamsResult{}
	_, err := manager.get(endpoints.teams.ListTeams, output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *teamsManager) EnableOps(data *teams.EnableOpsRequest) (*teams.SuccessResponse, error) {
	output := &teams.SuccessResponse{}
	err := manager.postJSON(endpoints.teams.EnableOps(data.TeamID), nil, output, nil, http.StatusAccepted, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *teamsManager) GetTeamRequestStatus(data *teams.GetTeamRequestStatusRequest) (*teams.RequestStatusResponse, error) {
	output := &teams.RequestStatusResponse{}
	_, err := manager.get(endpoints.teams.GetTeamRequestStatus(data.TeamID, data.RequestID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

