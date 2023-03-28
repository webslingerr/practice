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

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (c *userRepo) Create(ctx context.Context, req *models.CreateUser) (string, error) {
	var (
		query string
		id    = uuid.New().String()
	)

	query = `
		INSERT INTO users(
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

func (c *userRepo) GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {
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
		FROM users
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

	return &models.User{
		Id:        id.String,
		Name:      name.String,
		Phone:     phone.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}, nil
}

func (c *userRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error) {
	resp = &models.GetListUserResponse{}

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
		FROM users	
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

		var user models.User

		var id, name, phone, created_at, updated_at sql.NullString

		err = rows.Scan(
			&id,
			&name,
			&phone,
			&created_at,
			&updated_at,
		)

		user.Id = id.String
		user.Name = name.String
		user.Phone = phone.String
		user.CreatedAt = created_at.String
		user.UpdatedAt = updated_at.String

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &user)
	}

	resp.Count = len(resp.Users)

	return resp, nil
}

func (c *userRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			users
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

func (c *userRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			users
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

func (c *userRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) error {

	_, err := c.db.Exec(ctx,
		"DELETE FROM users WHERE id = $1", req.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
