package rotations

type Rotation struct {
	ID           string                   `json:"id,omitempty"`
	Name         string                   `json:"name,omitempty"`
	StartDate    string                   `json:"startDate,omitempty"`
	Length       int                      `json:"length,omitempty"`
	Participants []map[string]interface{} `json:"participants,omitempty"`
	Type         string                   `json:"type,omitempty"`
}

type ListRotationsResult struct {
	Links     RotationsResponseLink `json:"links"`
	Rotations []Rotation            `json:"values"`
	Count     int                   `json:"count"`
}

type RotationsResponseLink struct {
	Next string `json:"next"`
}

