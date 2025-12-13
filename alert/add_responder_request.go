package alert

type AddResponderRequest struct {
	ID        string
	Responder Responder `json:"responder"`
}

