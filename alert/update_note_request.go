package alert

type UpdateAlertNoteRequest struct {
	AlertID string
	NoteID  string
	Note    string `json:"note"`
}

