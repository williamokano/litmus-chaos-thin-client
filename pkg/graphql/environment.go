package graphql

import "github.com/williamokano/litmus-chaos-thin-client/pkg/entities"

type EnvironmentSortingField string

var (
	EnvironmentSortingFieldName EnvironmentSortingField = "NAME"
	EnvironmentSortingFieldTime EnvironmentSortingField = "TIME"
)

type CreateEnvironmentRequest struct {
	EnvironmentID string                   `json:"environmentID"`
	Name          string                   `json:"name"`
	Type          entities.EnvironmentType `json:"type"`
	Description   string                   `json:"description"`
	Tags          []string                 `json:"tags"`
}

type UpdateEnvironmentRequest struct {
	EnvironmentID string                   `json:"environmentID"`
	Name          string                   `json:"name"`
	Type          entities.EnvironmentType `json:"type"`
	Description   string                   `json:"description"`
	Tags          []string                 `json:"tags"`
}

type EnvironmentFilterInput struct {
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Type        string                     `json:"type"`
	Tags        []entities.EnvironmentType `json:"tags"`
	// Yes I know above types looks wrong, but they are wrongly defined
	// on their graphql schema.
	// please check this issue on their GH https://github.com/litmuschaos/litmus/issues/4354
}

type EnvironmentSortInput struct {
	Field     EnvironmentSortingField `json:"field"`
	Ascending bool                    `json:"ascending"`
}

type ListEnvironmentRequest struct {
	EnvironmentIDs []string                `json:"environmentIDs"`
	Pagination     *Pagination             `json:"pagination"`
	Filter         *EnvironmentFilterInput `json:"filter"`
	Sort           *EnvironmentSortInput   `json:"sort"`
}

type ListEnvironmentResponse struct {
	TotalNoOfEnvironments int                     `json:"totalNoOfEnvironments"`
	Environments          []*entities.Environment `json:"environments"`
}

type GetEnvironmentQuery struct {
	GetEnvironment struct {
		entities.Environment
	} `graphql:"getEnvironment(projectID: $projectID, environmentID: $environmentID)"`
}

type ListEnvironmentsQuery struct {
	ListEnvironments struct {
		ListEnvironmentResponse
	} `graphql:"listEnvironments(projectID: $projectID, request: $request)"`
}

type CreateEnvironmentMutation struct {
	CreateEnvironment struct {
		entities.Environment
	} `graphql:"createEnvironment(projectID: $projectID, request: $request)"`
}

type UpdateEnvironmentMutation struct {
	UpdateEnvironment struct {
		string
	} `graphql:"updateEnvironment(projectID: $projectID, request: $request)"`
}

type DeleteEnvironmentMutation struct {
	DeleteEnvironment struct {
		string
	} `graphql:"deleteEnvironment(projectID: $projectID, environmentID: $environmentID)"`
}
