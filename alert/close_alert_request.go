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

	switch r.IdentifierType {
	case ALIAS:
		query.Is("identifierType", "alias")
	case TINYID:
		query.Is("identifierType", "tiny")
	default:
		query.Is("identifierType", "id")
	}
	return query
}
