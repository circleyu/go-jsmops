package contacts

import (
	"strconv"

	"github.com/circleyu/go-jsmops/params"
)

type ListContactsRequest struct {
	Offset int
	Size   int
}

func (r *ListContactsRequest) RequestParams() *params.Params {
	query := params.Build()
	if r.Offset > 0 {
		query.Is("offset", strconv.Itoa(r.Offset))
	}
	if r.Size > 0 {
		query.Is("size", strconv.Itoa(r.Size))
	}
	return query
}

type CreateContactRequest struct {
	Method string `json:"method"`
	Value  string `json:"value"`
}

type GetContactRequest struct {
	ID string
}

type DeleteContactRequest struct {
	ID string
}

type UpdateContactRequest struct {
	ID    string
	Value string `json:"value"`
}

type ActivateContactRequest struct {
	ID string
}

type DeactivateContactRequest struct {
	ID string
}

