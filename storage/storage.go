package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Customer() CustomerRepoI
	User() UserRepoI
	Courier() CourierRepoI
	Category() CategoryRepoI
	Product() ProductRepoI
	Order() OrderRepoI
}

type CustomerRepoI interface {
	Create(context.Context, *models.CreateCustomer) (string, error)
	GetByID(context.Context, *models.CustomerPrimaryKey) (*models.Customer, error)
	GetList(context.Context, *models.GetListCustomerRequest) (*models.GetListCustomerResponse, error)
	Update(context.Context, *models.UpdateCustomer) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.CustomerPrimaryKey) error
}

type UserRepoI interface {
	Create(context.Context, *models.CreateUser) (string, error)
	GetByID(context.Context, *models.UserPrimaryKey) (*models.User, error)
	GetList(context.Context, *models.GetListUserRequest) (*models.GetListUserResponse, error)
	Update(context.Context, *models.UpdateUser) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.UserPrimaryKey) error
}

type CourierRepoI interface {
	Create(context.Context, *models.CreateCourier) (string, error)
	GetByID(context.Context, *models.CourierPrimaryKey) (*models.Courier, error)
	GetList(context.Context, *models.GetListCourierRequest) (*models.GetListCourierResponse, error)
	Update(context.Context, *models.UpdateCourier) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.CourierPrimaryKey) error
}

type CategoryRepoI interface {
	Create(context.Context, *models.CreateCategory) (string, error)
	GetByID(context.Context, *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(context.Context, *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Update(context.Context, *models.UpdateCategory) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.CategoryPrimaryKey) error
}

type ProductRepoI interface {
	Create(context.Context, *models.CreateProduct) (string, error)
	GetByID(context.Context, *models.ProductPrimaryKey) (*models.Product, error)
	GetList(context.Context, *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(context.Context, *models.UpdateProduct) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.ProductPrimaryKey) error
}

type OrderRepoI interface {
	Create(context.Context, *models.CreateOrder) (string, error)
	GetByID(context.Context, *models.OrderPrimaryKey) (*models.Order, error)
	GetList(context.Context, *models.GetListOrderRequest) (*models.GetListOrderResponse, error)
	Update(context.Context, *models.UpdateOrder) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.OrderPrimaryKey) error
}
