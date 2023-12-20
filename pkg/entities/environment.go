package entities

type EnvironmentType string

var (
	EnvironmentTypeProd    EnvironmentType = "PROD"
	EnvironmentTypeNonProd EnvironmentType = "NON_PROD"
)

type Environment struct {
	ResourceDetails `json:",inline"`
	Audit           `json:",inline"`
	ProjectID       string              `json:"projectID"`
	EnvironmentID   string              `json:"environmentID"`
	Name            string              `json:"name"`
	Description     string              `json:"description"`
	Tags            []string            `json:"tags"`
	Type            EnvironmentType     `json:"type"`
	InfraIDs        []string            `json:"infraIDs"`
	CreatedAt       int64               `json:"createdAt"`
	CreatedBy       *UserDetailResponse `json:"createdBy"`
	UpdatedAt       int64               `json:"updatedAt"`
	UpdatedBy       *UserDetailResponse `json:"updatedBy"`
	IsRemoved       bool                `json:"isRemoved"`
}
