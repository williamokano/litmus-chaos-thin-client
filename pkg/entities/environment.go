package entities

type EnvironmentType string

var (
	EnvironmentTypeProd    EnvironmentType = "PROD"
	EnvironmentTypeNonProd EnvironmentType = "NON_PROD"
)

type Environment struct {
	ResourceDetails
	Audit
	ProjectID     string              `graphql:"projectID" json:"projectID"`
	EnvironmentID string              `graphql:"environmentID" json:"environmentID"`
	Name          string              `graphql:"name" json:"name"`
	Description   string              `graphql:"description" json:"description"`
	Tags          []string            `graphql:"tags" json:"tags"`
	Type          EnvironmentType     `graphql:"type" json:"type"`
	InfraIDs      []string            `graphql:"infraIDs" json:"infraIDs"`
	CreatedAt     string              `graphql:"createdAt" json:"createdAt"`
	CreatedBy     *UserDetailResponse `graphql:"createdBy" json:"createdBy"`
	UpdatedAt     string              `graphql:"updatedAt" json:"updatedAt"`
	UpdatedBy     *UserDetailResponse `graphql:"updatedBy" json:"updatedBy"`
	IsRemoved     bool                `graphql:"isRemoved" json:"isRemoved"`
}
