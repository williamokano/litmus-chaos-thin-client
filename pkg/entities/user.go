package entities

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	// Audit fields are inconsistent leading to errors
	// when decoding. timestamp fields sometimes
	// come as int and sometimes as string
	// As audit is not important here, I'll just ignore
	//Audit         `json:",inline"`
	ID       string `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Name     string `json:"name,omitempty"`
	Role     Role   `json:"role"`
	//DeactivatedAt string `json:"deactivated_at,omitempty"`
}
