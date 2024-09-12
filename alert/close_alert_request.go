package alert

import (
	"github.com/circleyu/go-jsmops/params"
	"github.com/pkg/errors"
)

type CloseAlertRequest struct {
	IdentifierType  AlertIdentifier
	IdentifierValue string
}

func (r *CloseAlertRequest) Validate() error {
	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r *CloseAlertRequest) RequestParams() *params.Params {
	query := params.Build()

	if r.IdentifierType == ALIAS {
		query.Is("identifierType", "alias")

	} else if r.IdentifierType == TINYID {
		query.Is("identifierType", "tiny")

	} else {
		query.Is("identifierType", "id")

	}
	return query
}
