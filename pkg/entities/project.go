package entities

// Project contains the required fields to be stored in the database for a project
type Project struct {
	Audit `json:",inline"`
	ID    string  `graphql:"projectID" json:"projectID"`
	Name  string  `graphql:"name" json:"name"`
	State *string `graphql:"state" json:"state"`
}

type CreateProjectInput struct {
	ProjectName string `graphql:"projectName" json:"projectName"`
}

type UpdateProjectNameInput struct {
	ProjectID   string `graphql:"projectID" json:"projectID"`
	ProjectName string `graphql:"projectName" json:"projectName"`
}
