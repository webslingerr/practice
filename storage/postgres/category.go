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

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (c *categoryRepo) Create(ctx context.Context, req *models.CreateCategory) (string, error) {
	var (
		query string
		id    = uuid.New().String()
	)

	query = `
		INSERT INTO categories(
			id, 
			name,
			updated_at
		)
		VALUES (:id, :name, NOW())
	`

	params := map[string]interface{}{
		"id":    id,
		"name":  req.Name,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	_, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *categoryRepo) GetByID(ctx context.Context, req *models.CategoryPrimaryKey) (*models.Category, error) {
	var (
		query      string
		id         sql.NullString
		name       sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)

	query = `
		SELECT 
			id,
			name,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24-MI-SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24-MI-SS')
		FROM categories
		WHERE id = $1
	`

	err := c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&created_at,
		&updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &models.Category{
		Id:        id.String,
		Name:      name.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}, nil
}

func (c *categoryRepo) GetList(ctx context.Context, req *models.GetListCategoryRequest) (resp *models.GetListCategoryResponse, err error) {
	resp = &models.GetListCategoryResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			id, 
			name,
			created_at,
			updated_at
		FROM categories	
	`

	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
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

		var courier models.Category

		var id, name, created_at, updated_at sql.NullString

		err = rows.Scan(
			&id,
			&name,
			&created_at,
			&updated_at,
		)

		courier.Id = id.String
		courier.Name = name.String
		courier.CreatedAt = created_at.String
		courier.UpdatedAt = updated_at.String

		if err != nil {
			return nil, err
		}

		resp.Categories = append(resp.Categories, &courier)
	}

	resp.Count = len(resp.Categories)

	return resp, nil
}

func (c *categoryRepo) Update(ctx context.Context, req *models.UpdateCategory) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			categories
		SET 
			name = :name,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":    req.Id,
		"name":  req.Name,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (c *categoryRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			categories
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

func (c *categoryRepo) Delete(ctx context.Context, req *models.CategoryPrimaryKey) error {

	_, err := c.db.Exec(ctx,
		"DELETE FROM categories WHERE id = $1", req.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
