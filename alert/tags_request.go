package alert

type AddTagsRequest struct {
	ID   string
	Tags []string `json:"tags"`
}

type DeleteTagsRequest struct {
	ID   string
	Tags []string `json:"tags"`
}

