package postgres

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *orderRepo {
	return &orderRepo{
		db: db,
	}
}

func (o *orderRepo) Create(ctx context.Context, req *models.CreateOrder) (string, error) {
	var (
		query string
		id    = uuid.New().String()
	)

	query = `
		INSERT INTO orders(
			id,
			name,
			quantity,
			user_id,
			customer_id,
			product_id,
			courier_id,
			updated_at
		) VALUES
		(:id, :name, :quantity, :user_id, :customer_id, :product_id, :courier_id, NOW())
	`

	params := map[string]interface{}{
		"id":          id,
		"name":        req.Name,
		"quantity":    req.Quantity,
		"user_id":     helper.NewNullString(req.UserId),
		"customer_id": helper.NewNullString(req.CustomerId),
		"product_id":  helper.NewNullString(req.ProductId),
		"courier_id":  helper.NewNullString(req.CourierId),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	_, err := o.db.Exec(ctx, query, args...)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (o *orderRepo) GetByID(ctx context.Context, req *models.OrderPrimaryKey) (*models.Order, error) {
	var (
		query          string
		id             sql.NullString
		name           sql.NullString
		product_price  sql.NullFloat64
		total_price    sql.NullFloat64
		quantity       sql.NullInt32
		user_name      sql.NullString
		user_phone     sql.NullString
		product_name   sql.NullString
		customer_name  sql.NullString
		customer_phone sql.NullString
		courier_name   sql.NullString
		courier_phone  sql.NullString
		created_at     sql.NullString
		updated_at     sql.NullString
	)

	query = `
		SELECT 
			o.id,
			o.name,
			o.price,
			total_price,
			quantity,
			u.name,
			u.phone,
			p.name,
			c.name,
			c.phone,
			co.name,
			co.phone,
			o.created_at,
			o.updated_at
		FROM orders AS o
		LEFT JOIN users AS u ON o.user_id = u.id
		LEFT JOIN customers AS c ON o.customer_id = c.id
		LEFT JOIN products AS p ON o.product_id = p.id
		LEFT JOIN couriers AS co ON o.courier_id = co.id
		WHERE o.id = $1
	`

	err := o.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&product_price,
		&total_price,
		&quantity,
		&user_name,
		&user_phone,
		&product_name,
		&customer_name,
		&customer_phone,
		&courier_name,
		&courier_phone,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}

	var user models.ReturnUser
	user.Name = user_name.String
	user.Phone = user_phone.String

	var product models.ReturnProduct
	product.Name = product_name.String
	product.Price = product_price.Float64

	var customer models.ReturnCustomer
	customer.Name = customer_name.String
	customer.Phone = customer_phone.String

	return &models.Order{
		Id:         id.String,
		Name:       name.String,
		Price:      product_price.Float64,
		TotalPrice: total_price.Float64,
		Quantity:   quantity.Int32,
		User:       user,
		Customer:   customer,
		Product:    product,
		CreatedAt:  created_at.String,
		UpdatedAt:  updated_at.String,
	}, nil
}

func (o *orderRepo) GetList(ctx context.Context, req *models.GetListOrderRequest) (resp *models.GetListOrderResponse, err error) {
	resp = &models.GetListOrderResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT 
			o.id,
			o.name,
			o.price,
			total_price,
			quantity,
			u.name,
			u.phone,
			p.name,
			c.name,
			c.phone,
			co.name,
			co.phone,
			o.created_at,
			o.updated_at
		FROM orders AS o
		LEFT JOIN users AS u ON o.user_id = u.id
		LEFT JOIN customers AS c ON o.customer_id = c.id
		LEFT JOIN products AS p ON o.product_id = p.id
		LEFT JOIN couriers AS co ON o.courier_id = co.id
	`

	if len(req.Search) > 0 {
		filter += " AND o.name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := o.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var order models.Order
		var user models.ReturnUser
		var customer models.ReturnCustomer
		var courier models.ReturnCourier
		var product models.ReturnProduct

		var (
			id             sql.NullString
			name           sql.NullString
			user_name      sql.NullString
			user_phone     sql.NullString
			courier_name   sql.NullString
			courier_phone  sql.NullString
			customer_name  sql.NullString
			customer_phone sql.NullString
			product_name   sql.NullString
			created_at     sql.NullString
			updated_at     sql.NullString

			price       sql.NullFloat64
			total_price sql.NullFloat64

			quantity sql.NullInt32
		)

		err = rows.Scan(
			&id,
			&name,
			&price,
			&total_price,
			&quantity,
			&user_name,
			&user_phone,
			&product_name,
			&customer_name,
			&customer_phone,
			&courier_name,
			&courier_phone,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}

		user.Name = user_name.String
		user.Phone = user_phone.String
		courier.Name = courier_name.String
		courier.Phone = courier_phone.String
		customer.Name = customer_name.String
		customer.Phone = customer_phone.String
		product.Name = product_name.String
		product.Price = price.Float64

		order.User = user
		order.Product = product
		order.Customer = customer
		order.Courier = courier

		resp.Orders = append(resp.Orders, &order)
	}

	resp.Count = len(resp.Orders)

	return resp, nil
}

func (o *orderRepo) Update(ctx context.Context, req *models.UpdateOrder) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			orders
		SET 
			name = :name,
			quantity = :quantity,
			user_id = :user_id,
			product_id = :product_id,
			customer_id = :customer_id,
			courier_id = :courier_id,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":          req.Id,
		"name":        req.Name,
		"quantity":    req.Quantity,
		"user_id":     helper.NewNullString(req.UserId),
		"product_id":  helper.NewNullString(req.ProductId),
		"customer_id": helper.NewNullString(req.CustomerId),
		"courier_id":  helper.NewNullString(req.CourierId),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := o.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (o *orderRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

	var (
		query string
		set   string
	)

	if len(req.Fields) <= 0 {
		return 0, errors.New("no fields")
	}

	for key := range req.Fields {
		set += fmt.Sprintf(" %s = :%s, ", key, key)
	}

	query = `
		UPDATE
			orders
		SET
	` + set + ` updated_at = now()
		WHERE id = :id
	`

	req.Fields["id"] = req.ID

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	result, err := o.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (o *orderRepo) Delete(ctx context.Context, req *models.OrderPrimaryKey) error {
	_, err := o.db.Exec(ctx,
		"DELETE FROM products WHERE id = $1", req.Id,
	)

	if err != nil {
		return err
	}

	return nil
}