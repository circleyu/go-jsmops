package contacts

type Contact struct {
	ID       string `json:"id,omitempty"`
	Method   string `json:"method,omitempty"`
	Value    string `json:"value,omitempty"`
	Enabled  bool   `json:"enabled,omitempty"`
}

type ListContactsResult struct {
	Links    ContactsResponseLink `json:"links"`
	Contacts []Contact            `json:"values"`
	Count    int                  `json:"count"`
}

type ContactsResponseLink struct {
	Next string `json:"next"`
}

