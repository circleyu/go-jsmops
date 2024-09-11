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

	if r.IdentifierType == ALIAS {
		params["identifierType"] = "alias"

	} else if r.IdentifierType == TINYID {
		params["identifierType"] = "tiny"

	} else {
		params["identifierType"] = "id"

	}
	return params
}
