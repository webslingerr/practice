package models

type Courier struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ReturnCourier struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type CourierPrimaryKey struct {
	Id string `json:"id"`
}

type CreateCourier struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type UpdateCourier struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	UpdatedAt string `json:"updated_at"`
}

type GetListCourierRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListCourierResponse struct {
	Count    int        `json:"count"`
	Couriers []*Courier `json:"couriers"`
}
