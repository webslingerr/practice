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

type courierRepo struct {
	db *pgxpool.Pool
}

func NewCourierRepo(db *pgxpool.Pool) *courierRepo {
	return &courierRepo{
		db: db,
	}
}

func (c *courierRepo) Create(ctx context.Context, req *models.CreateCourier) (string, error) {
	var (
		query string
		id    = uuid.New().String()
	)

	query = `
		INSERT INTO couriers(
			id, 
			name,
			phone,
			updated_at
		)
		VALUES (:id, :name, :phone, NOW())
	`

	params := map[string]interface{}{
		"id":    id,
		"name":  req.Name,
		"phone": req.Phone,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	_, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *courierRepo) GetByID(ctx context.Context, req *models.CourierPrimaryKey) (*models.Courier, error) {
	var (
		query      string
		id         sql.NullString
		name       sql.NullString
		phone      sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)

	query = `
		SELECT 
			id,
			name,
			phone,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24-MI-SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24-MI-SS')
		FROM couriers
		WHERE id = $1
	`

	err := c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&phone,
		&created_at,
		&updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &models.Courier{
		Id:        id.String,
		Name:      name.String,
		Phone:     phone.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}, nil
}

func (c *courierRepo) GetList(ctx context.Context, req *models.GetListCourierRequest) (resp *models.GetListCourierResponse, err error) {
	resp = &models.GetListCourierResponse{}

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
			phone,
			created_at,
			updated_at
		FROM couriers	
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

		var courier models.Courier

		var id, name, phone, created_at, updated_at sql.NullString

		err = rows.Scan(
			&id,
			&name,
			&phone,
			&created_at,
			&updated_at,
		)

		courier.Id = id.String
		courier.Name = name.String
		courier.Phone = phone.String
		courier.CreatedAt = created_at.String
		courier.UpdatedAt = updated_at.String

		if err != nil {
			return nil, err
		}

		resp.Couriers = append(resp.Couriers, &courier)
	}

	resp.Count = len(resp.Couriers)

	return resp, nil
}

func (c *courierRepo) Update(ctx context.Context, req *models.UpdateCourier) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			couriers
		SET 
			name = :name,
			phone = :phone,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":    req.Id,
		"name":  req.Name,
		"phone": req.Phone,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (c *courierRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			couriers
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

func (c *courierRepo) Delete(ctx context.Context, req *models.CourierPrimaryKey) error {

	_, err := c.db.Exec(ctx,
		"DELETE FROM couriers WHERE id = $1", req.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
