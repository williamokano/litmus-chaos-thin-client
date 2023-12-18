package entities

// Project contains the required fields to be stored in the database for a project
type Project struct {
	Audit `json:",inline"`
	ID    string  `json:"projectID"`
	Name  string  `json:"name"`
	State *string `json:"state"`
}

type CreateProjectInput struct {
	ProjectName string `json:"projectName"`
}

type UpdateProjectNameInput struct {
	ProjectID   string `json:"projectID"`
	ProjectName string `json:"projectName"`
}
