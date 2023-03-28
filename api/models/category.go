package models

type Category struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ReturnCategory struct {
	Name string `json:"name"`
}

type CategoryPrimaryKey struct {
	Id string `json:"id"`
}

type CreateCategory struct {
	Name string `json:"name"`
}

type UpdateCategory struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetListCategoryRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListCategoryResponse struct {
	Count      int         `json:"count"`
	Categories []*Category `json:"categories"`
}
