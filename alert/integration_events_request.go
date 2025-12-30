package alert

import (
	"github.com/circleyu/go-jsmops/v2/params"
	"github.com/pkg/errors"
)

// IntegrationCreateAlertRequest represents a request to create an alert via Integration Events API
// This uses 'details' instead of 'extraProperties' and includes 'user' and 'note' fields
type IntegrationCreateAlertRequest struct {
	Message     string            `json:"message"`
	Alias       string            `json:"alias,omitempty"`
	Description string            `json:"description,omitempty"`
	Responders  []Responder       `json:"responders,omitempty"`
	VisibleTo   []Responder       `json:"visibleTo,omitempty"`
	Actions     []string          `json:"actions,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	Details     map[string]string `json:"details,omitempty"` // Note: uses 'details' instead of 'extraProperties'
	Entity      string            `json:"entity,omitempty"`
	Source      string            `json:"source,omitempty"`
	Priority    Priority          `json:"priority,omitempty"`
	User        string            `json:"user,omitempty"`
	Note        string            `json:"note,omitempty"`
}

func (r *IntegrationCreateAlertRequest) Validate() error {
	if r.Message == "" {
		return errors.New("message can not be empty")
	}
	return nil
}

// IntegrationAcknowledgeAlertRequest represents a request to acknowledge an alert via Integration Events API
type IntegrationAcknowledgeAlertRequest struct {
	IdentifierType  AlertIdentifier
	IdentifierValue string
	User            string `json:"user,omitempty"`
	Source          string `json:"source,omitempty"`
	Note            string `json:"note,omitempty"`
}

func (r *IntegrationAcknowledgeAlertRequest) Validate() error {
	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r *IntegrationAcknowledgeAlertRequest) RequestParams() *params.Params {
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

// IntegrationCloseAlertRequest represents a request to close an alert via Integration Events API
type IntegrationCloseAlertRequest struct {
	IdentifierType  AlertIdentifier
	IdentifierValue string
	User            string `json:"user,omitempty"`
	Source          string `json:"source,omitempty"`
	Note            string `json:"note,omitempty"`
}

func (r *IntegrationCloseAlertRequest) Validate() error {
	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r *IntegrationCloseAlertRequest) RequestParams() *params.Params {
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

// IntegrationAddNoteRequest represents a request to add a note to an alert via Integration Events API
type IntegrationAddNoteRequest struct {
	IdentifierType  AlertIdentifier
	IdentifierValue string
	User            string `json:"user,omitempty"`
	Source          string `json:"source,omitempty"`
	Note            string `json:"note,omitempty"`
}

func (r *IntegrationAddNoteRequest) Validate() error {
	if r.Note == "" {
		return errors.New("Note can not be empty")
	}
	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r *IntegrationAddNoteRequest) RequestParams() *params.Params {
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

