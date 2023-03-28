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

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (c *productRepo) Create(ctx context.Context, req *models.CreateProduct) (string, error) {
	var (
		query string
		id    = uuid.New().String()
	)

	query = `
		INSERT INTO products(
			id, 
			name,
			price,
			category_id,
			updated_at
		)
		VALUES (:id, :name, :price, :category_id, NOW())
	`

	params := map[string]interface{}{
		"id":          id,
		"name":        req.Name,
		"price":       req.Price,
		"category_id": helper.NewNullString(req.CategoryId),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	_, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *productRepo) GetByID(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error) {
	var (
		query         string
		id            sql.NullString
		name          sql.NullString
		price         sql.NullFloat64
		category_name sql.NullString
		created_at    sql.NullString
		updated_at    sql.NullString
	)

	query = `
		SELECT 
			p.id,
			p.name,
			price,
			COALESCE(c.name, ''),
			TO_CHAR(p.created_at, 'YYYY-MM-DD HH24-MI-SS'),
			TO_CHAR(p.updated_at, 'YYYY-MM-DD HH24-MI-SS')
		FROM products AS p
		LEFT JOIN categories AS c ON p.category_id = c.id
		WHERE p.id = $1
	`

	err := c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&price,
		&category_name,
		&created_at,
		&updated_at,
	)

	var category models.ReturnCategory
	category.Name = category_name.String

	if err != nil {
		return nil, err
	}

	return &models.Product{
		Id:        id.String,
		Name:      name.String,
		Price:     price.Float64,
		Category:  category,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}, nil
}

func (c *productRepo) GetList(ctx context.Context, req *models.GetListProductRequest) (resp *models.GetListProductResponse, err error) {
	resp = &models.GetListProductResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			p.id, 
			p.name,
			p.price,
			c.name,
			p.created_at,
			p.updated_at
		FROM products AS p
		LEFT JOIN categories AS c ON p.category_id = c.id
	`

	if len(req.Search) > 0 {
		filter += " AND p.name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := c.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var product models.Product
		var category models.ReturnCategory

		var id, name, category_name, created_at, updated_at sql.NullString
		var price sql.NullFloat64

		err = rows.Scan(
			&id,
			&name,
			&price.Float64,
			&category_name,
			&created_at,
			&updated_at,
		)

		product.Id = id.String
		product.Name = name.String
		category.Name = category_name.String
		product.Price = price.Float64
		product.CreatedAt = created_at.String
		product.UpdatedAt = updated_at.String

		if err != nil {
			return nil, err
		}

		product.Category = category
		resp.Products = append(resp.Products, &product)
	}

	resp.Count = len(resp.Products)

	return resp, nil
}

func (c *productRepo) Update(ctx context.Context, req *models.UpdateProduct) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			products
		SET 
			name = :name,
			price = :price,
			category_id = :category_id,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":          req.Id,
		"name":        req.Name,
		"price":       req.Price,
		"category_id": req.CategoryId,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (c *productRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			products
		SET
	` + set + ` updated_at = now()
		WHERE id = :id
	`

	req.Fields["id"] = req.ID

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (c *productRepo) Delete(ctx context.Context, req *models.ProductPrimaryKey) error {

	_, err := c.db.Exec(ctx,
		"DELETE FROM products WHERE id = $1", req.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
