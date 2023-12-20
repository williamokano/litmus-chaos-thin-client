package graphql

import (
	"strings"

	"github.com/williamokano/litmus-chaos-thin-client/pkg/entities"
)

type EnvironmentSortingField string

var (
	EnvironmentSortingFieldName EnvironmentSortingField = "NAME"
	EnvironmentSortingFieldTime EnvironmentSortingField = "TIME"
)

type CreateEnvironmentRequest struct {
	EnvironmentID string                   `graphql:"environmentID" json:"environmentID"`
	Name          string                   `graphql:"name" json:"name"`
	Type          entities.EnvironmentType `graphql:"type" json:"type"`
	Description   string                   `graphql:"description" json:"description"`
	Tags          []string                 `graphql:"tags" json:"tags"`
}

// IDFromName tries to mimic the Control Plane UI behaviour
func (r CreateEnvironmentRequest) IDFromName() string {
	return strings.NewReplacer("-", "", "_", "", " ", "_").Replace(r.Name)
}

type UpdateEnvironmentRequest struct {
	EnvironmentID string                   `graphql:"environmentID" json:"environmentID"`
	Name          string                   `graphql:"name" json:"name"`
	Type          entities.EnvironmentType `graphql:"type" json:"type"`
	Description   string                   `graphql:"description" json:"description"`
	Tags          []string                 `graphql:"tags" json:"tags"`
}

type EnvironmentFilterInput struct {
	Name        string                     `graphql:"name" json:"name"`
	Description string                     `graphql:"description" json:"description"`
	Type        string                     `graphql:"type" json:"type"`
	Tags        []entities.EnvironmentType `graphql:"tags" json:"tags"`
	// Yes I know above types looks wrong, but they are wrongly defined
	// on their graphql schema.
	// please check this issue on their GH https://github.com/litmuschaos/litmus/issues/4354
}

type EnvironmentSortInput struct {
	Field     EnvironmentSortingField `graphql:"field" json:"field"`
	Ascending bool                    `graphql:"ascending" json:"ascending"`
}

type ListEnvironmentRequest struct {
	EnvironmentIDs []string                `graphql:"environmentIDs" json:"environmentIDs"`
	Pagination     *Pagination             `graphql:"pagination" json:"pagination"`
	Filter         *EnvironmentFilterInput `graphql:"filter" json:"filter"`
	Sort           *EnvironmentSortInput   `graphql:"sort" json:"sort"`
}

type ListEnvironmentResponse struct {
	TotalNoOfEnvironments int                     `graphql:"totalNoOfEnvironments" json:"totalNoOfEnvironments"`
	Environments          []*entities.Environment `graphql:"environments" json:"environments"`
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
