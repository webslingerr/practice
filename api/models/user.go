package models

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ReturnUser struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type CreateUser struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type UpdateUser struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	UpdatedAt string `json:"updated_at"`
}

type GetListUserRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListUserResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"users"`
}
