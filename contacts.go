package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/v2/contacts"
)

type ContactsManager interface {
	ListContacts(*contacts.ListContactsRequest) (*contacts.ListContactsResult, error)
	CreateContact(*contacts.CreateContactRequest) (*contacts.Contact, error)
	GetContact(*contacts.GetContactRequest) (*contacts.Contact, error)
	DeleteContact(*contacts.DeleteContactRequest) error
	UpdateContact(*contacts.UpdateContactRequest) (*contacts.Contact, error)
	ActivateContact(*contacts.ActivateContactRequest) (*contacts.Contact, error)
	DeactivateContact(*contacts.DeactivateContactRequest) (*contacts.Contact, error)
}

type contactsManager struct {
	*APIClient
}

func newContactsManager(client *APIClient) *contactsManager {
	return &contactsManager{
		client,
	}
}

func (manager *contactsManager) ListContacts(data *contacts.ListContactsRequest) (*contacts.ListContactsResult, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &contacts.ListContactsResult{}
	_, err := manager.get(endpoints.contacts.ListContacts, output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *contactsManager) CreateContact(data *contacts.CreateContactRequest) (*contacts.Contact, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &contacts.Contact{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.contacts.CreateContact, jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *contactsManager) GetContact(data *contacts.GetContactRequest) (*contacts.Contact, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &contacts.Contact{}
	_, err := manager.get(endpoints.contacts.GetContact(data.ID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *contactsManager) DeleteContact(data *contacts.DeleteContactRequest) error {
	if err := manager.checkBasicAuth(); err != nil {
		return err
	}
	return manager.delete(endpoints.contacts.DeleteContact(data.ID), nil, http.StatusNoContent, http.StatusOK)
}

func (manager *contactsManager) UpdateContact(data *contacts.UpdateContactRequest) (*contacts.Contact, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &contacts.Contact{}
	requestBody := map[string]interface{}{
		"value": data.Value,
	}
	jsonb, err := sonic.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	err = manager.patch(endpoints.contacts.UpdateContact(data.ID), jsonb, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *contactsManager) ActivateContact(data *contacts.ActivateContactRequest) (*contacts.Contact, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &contacts.Contact{}
	err := manager.patch(endpoints.contacts.ActivateContact(data.ID), nil, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *contactsManager) DeactivateContact(data *contacts.DeactivateContactRequest) (*contacts.Contact, error) {
	if err := manager.checkBasicAuth(); err != nil {
		return nil, err
	}
	output := &contacts.Contact{}
	err := manager.patch(endpoints.contacts.DeactivateContact(data.ID), nil, output, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}
