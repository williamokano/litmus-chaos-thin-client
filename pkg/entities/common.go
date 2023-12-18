package entities

type ResourceDetails struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

type Audit struct {
	UpdatedAt int64              `json:"updatedAt"`
	CreatedAt int64              `json:"createdAt"`
	CreatedBy UserDetailResponse `json:"createdBy"`
	UpdatedBy UserDetailResponse `json:"updatedBy"`
	IsRemoved bool               `json:"isRemoved"`
}

type UserDetailResponse struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
