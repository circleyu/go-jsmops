package alert

import (
	"github.com/pkg/errors"
)

type AddNoteRequest struct {
	IdentifierType  AlertIdentifier
	IdentifierValue string
	Note            string `json:"note"`
}

func (r *AddNoteRequest) Validate() error {
	if r.Note == "" {
		return errors.New("Note can not be empty")
	}
	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r *AddNoteRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	switch r.IdentifierType {
	case ALIAS:
		params["identifierType"] = "alias"
	case TINYID:
		params["identifierType"] = "tiny"
	default:
		params["identifierType"] = "id"
	}
	return params
}
