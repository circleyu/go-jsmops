package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/alert"
)

type AlertsManager interface {
	GetRequestStatus(*alert.GetRequestStatusRequest) (*alert.RequestStatusResponse, error)
	CreateAlert(*alert.CreateAlertRequest) (*alert.SuccessResponse, error)
	ListAlerts(*alert.ListAlertsRequest) (*alert.ListAlertsResult, error)
	GetAlert(*alert.GetAlertRequest) (*alert.Alert, error)
	DeleteAlert(*alert.DeleteAlertRequest) error
	GetAlertByAlias(*alert.GetAlertByAliasRequest) (*alert.Alert, error)
	AcknowledgeAlert(*alert.AcknowledgeAlertRequest) (*alert.SuccessResponse, error)
	AssignAlert(*alert.AssignAlertRequest) (*alert.SuccessResponse, error)
	AddResponder(*alert.AddResponderRequest) (*alert.SuccessResponse, error)
	AddExtraProperties(*alert.AddExtraPropertiesRequest) (*alert.SuccessResponse, error)
	DeleteExtraProperties(*alert.DeleteExtraPropertiesRequest) error
	AddTags(*alert.AddTagsRequest) (*alert.SuccessResponse, error)
	DeleteTags(*alert.DeleteTagsRequest) error
	CloseAlert(*alert.CloseAlertRequest) (*alert.SuccessResponse, error)
	EscalateAlert(*alert.EscalateAlertRequest) (*alert.SuccessResponse, error)
	ExecuteCustomAction(*alert.ExecuteCustomActionRequest) (*alert.SuccessResponse, error)
	SnoozeAlert(*alert.SnoozeAlertRequest) (*alert.SuccessResponse, error)
	UnacknowledgeAlert(*alert.UnacknowledgeAlertRequest) (*alert.SuccessResponse, error)
	ListAlertNotes(*alert.ListAlertNotesRequest) (*alert.ListAlertNotesResult, error)
	AddAlertNote(*alert.AddNoteRequest) (*alert.AddNoteResponse, error)
	DeleteAlertNote(*alert.DeleteAlertNoteRequest) error
	UpdateAlertNote(*alert.UpdateAlertNoteRequest) (*alert.AlertNote, error)
	UpdateAlertPriority(*alert.UpdateAlertPriorityRequest) (*alert.SuccessResponse, error)
	UpdateAlertMessage(*alert.UpdateAlertMessageRequest) (*alert.SuccessResponse, error)
	UpdateAlertDescription(*alert.UpdateAlertDescriptionRequest) (*alert.SuccessResponse, error)
	ListAlertLogs(*alert.ListAlertLogsRequest) (*alert.ListAlertLogsResult, error)
}

type alertsManager struct {
	*APIClient
}

func newAlertsManager(client *APIClient) *alertsManager {
	return &alertsManager{
		client,
	}
}

func (manager *alertsManager) CreateAlert(data *alert.CreateAlertRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.alerts.CreateAlert, jsonb, output, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (manager *alertsManager) AcknowledgeAlert(data *alert.AcknowledgeAlertRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	err := manager.post(endpoints.alerts.AcknowledgeAlert(data.IdentifierValue), nil, output, nil, http.StatusAccepted)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (manager *alertsManager) CloseAlert(data *alert.CloseAlertRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	err := manager.post(endpoints.alerts.CloseAlert(data.IdentifierValue), nil, output, nil, http.StatusAccepted)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (manager *alertsManager) GetRequestStatus(data *alert.GetRequestStatusRequest) (*alert.RequestStatusResponse, error) {
	output := &alert.RequestStatusResponse{}
	_, err := manager.get(endpoints.alerts.GetRequestStatus(data.RequestID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) GetAlert(data *alert.GetAlertRequest) (*alert.Alert, error) {
	output := &alert.Alert{}
	_, err := manager.get(endpoints.alerts.GetAlert(data.ID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) DeleteAlert(data *alert.DeleteAlertRequest) error {
	return manager.delete(endpoints.alerts.DeleteAlert(data.ID), nil, http.StatusNoContent, http.StatusOK)
}

func (manager *alertsManager) GetAlertByAlias(data *alert.GetAlertByAliasRequest) (*alert.Alert, error) {
	output := &alert.Alert{}
	_, err := manager.get(endpoints.alerts.GetAlertByAlias, output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) AssignAlert(data *alert.AssignAlertRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	requestBody := make(map[string]interface{})
	if data.Owner != nil {
		requestBody["owner"] = data.Owner
	}
	if data.OwnerTeam != "" {
		requestBody["ownerTeam"] = data.OwnerTeam
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.alerts.AssignAlert(data.ID), jsonb, output, nil, http.StatusAccepted)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) AddResponder(data *alert.AddResponderRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	requestBody := map[string]interface{}{
		"responder": data.Responder,
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.alerts.AddResponder(data.ID), jsonb, output, nil, http.StatusAccepted)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) AddExtraProperties(data *alert.AddExtraPropertiesRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	requestBody := map[string]interface{}{
		"properties": data.Properties,
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.alerts.AddExtraProperties(data.ID), jsonb, output, nil, http.StatusAccepted)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) DeleteExtraProperties(data *alert.DeleteExtraPropertiesRequest) error {
	requestBody := map[string]interface{}{
		"properties": data.Properties,
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return err
	}
	return manager.delete(endpoints.alerts.DeleteExtraProperties(data.ID), jsonb, http.StatusNoContent, http.StatusOK)
}

func (manager *alertsManager) AddTags(data *alert.AddTagsRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	requestBody := map[string]interface{}{
		"tags": data.Tags,
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.alerts.AddTags(data.ID), jsonb, output, nil, http.StatusAccepted)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) DeleteTags(data *alert.DeleteTagsRequest) error {
	requestBody := map[string]interface{}{
		"tags": data.Tags,
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return err
	}
	return manager.delete(endpoints.alerts.DeleteTags(data.ID), jsonb, http.StatusNoContent, http.StatusOK)
}

func (manager *alertsManager) EscalateAlert(data *alert.EscalateAlertRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	requestBody := make(map[string]interface{})
	if data.Escalation.ID != "" || data.Escalation.Name != "" {
		requestBody["escalation"] = data.Escalation
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.alerts.EscalateAlert(data.ID), jsonb, output, nil, http.StatusAccepted)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) ExecuteCustomAction(data *alert.ExecuteCustomActionRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	requestBody := map[string]interface{}{
		"action": data.Action,
	}
	if data.User != "" {
		requestBody["user"] = data.User
	}
	if data.Note != "" {
		requestBody["note"] = data.Note
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.alerts.ExecuteCustomAction(data.ID), jsonb, output, nil, http.StatusAccepted)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) SnoozeAlert(data *alert.SnoozeAlertRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	requestBody := make(map[string]interface{})
	if data.Until != nil {
		requestBody["until"] = data.Until
	}
	if data.UntilTime != "" {
		requestBody["untilTime"] = data.UntilTime
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.alerts.SnoozeAlert(data.ID), jsonb, output, nil, http.StatusAccepted)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) UnacknowledgeAlert(data *alert.UnacknowledgeAlertRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	err := manager.post(endpoints.alerts.UnacknowledgeAlert(data.ID), nil, output, nil, http.StatusAccepted)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) ListAlertNotes(data *alert.ListAlertNotesRequest) (*alert.ListAlertNotesResult, error) {
	output := &alert.ListAlertNotesResult{}
	_, err := manager.get(endpoints.alerts.ListAlertNotes(data.ID), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) AddAlertNote(data *alert.AddNoteRequest) (*alert.AddNoteResponse, error) {
	output := &alert.AddNoteResponse{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.alerts.AddAlertNote(data.IdentifierValue), jsonb, output, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) DeleteAlertNote(data *alert.DeleteAlertNoteRequest) error {
	return manager.delete(endpoints.alerts.DeleteAlertNote(data.AlertID, data.NoteID), nil, http.StatusNoContent, http.StatusOK)
}

func (manager *alertsManager) UpdateAlertNote(data *alert.UpdateAlertNoteRequest) (*alert.AlertNote, error) {
	output := &alert.AlertNote{}
	requestBody := map[string]interface{}{
		"note": data.Note,
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.alerts.UpdateAlertNote(data.AlertID, data.NoteID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) UpdateAlertPriority(data *alert.UpdateAlertPriorityRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.alerts.UpdateAlertPriority(data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) UpdateAlertMessage(data *alert.UpdateAlertMessageRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.alerts.UpdateAlertMessage(data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) UpdateAlertDescription(data *alert.UpdateAlertDescriptionRequest) (*alert.SuccessResponse, error) {
	output := &alert.SuccessResponse{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.alerts.UpdateAlertDescription(data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) ListAlertLogs(data *alert.ListAlertLogsRequest) (*alert.ListAlertLogsResult, error) {
	output := &alert.ListAlertLogsResult{}
	_, err := manager.get(endpoints.alerts.ListAlertLogs(data.ID), output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *alertsManager) ListAlerts(data *alert.ListAlertsRequest) (*alert.ListAlertsResult, error) {
	output := &alert.ListAlertsResult{}

	_, err := manager.get(endpoints.alerts.ListAlerts, output, data.RequestParams())
	if err != nil {
		return nil, err
	}

	return output, nil
}
