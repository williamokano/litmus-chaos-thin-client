package graphql

type Pagination struct {
	Page  int `graphql:"page" json:"page"`
	Limit int `graphql:"limit" json:"limit"`
}
