package entities

type ResourceDetails struct {
	Name        string   `graphql:"name" json:"name"`
	Description string   `graphql:"description" json:"description"`
	Tags        []string `graphql:"tags" json:"tags"`
}

type Audit struct {
	UpdatedAt string             `graphql:"updatedAt" json:"updatedAt"`
	CreatedAt string             `graphql:"createdAt" json:"createdAt"`
	CreatedBy UserDetailResponse `graphql:"createdBy" json:"createdBy"`
	UpdatedBy UserDetailResponse `graphql:"updatedBy" json:"updatedBy"`
	IsRemoved bool               `graphql:"isRemoved" json:"isRemoved"`
}

type UserDetailResponse struct {
	UserID   string `graphql:"userID" json:"userID"`
	Username string `graphql:"username" json:"username"`
	Email    string `graphql:"email" json:"email"`
}
