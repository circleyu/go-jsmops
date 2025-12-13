package alert

import (
	"github.com/pkg/errors"
)

type AcknowledgeAlertRequest struct {
	IdentifierType  AlertIdentifier
	IdentifierValue string
}

func (r *AcknowledgeAlertRequest) Validate() error {
	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r *AcknowledgeAlertRequest) RequestParams() map[string]string {

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
