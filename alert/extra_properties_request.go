package alert

type AddExtraPropertiesRequest struct {
	ID         string
	Properties map[string]string `json:"properties"`
}

type DeleteExtraPropertiesRequest struct {
	ID         string
	Properties []string `json:"properties"`
}

