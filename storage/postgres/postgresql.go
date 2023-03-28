package postgres

import (
	"app/config"
	"app/storage"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db *pgxpool.Pool
	customer storage.CustomerRepoI
	user storage.UserRepoI
	courier storage.CourierRepoI
	category storage.CategoryRepoI
	product storage.ProductRepoI
	order storage.OrderRepoI
}

func NewConnectPostgresql(cfg *config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))
	if err != nil {
		return nil, err
	}

	pgpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	if err := pgpool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &Store{
		db: pgpool,
		customer: NewCustomerRepo(pgpool),
		user: NewUserRepo(pgpool),
		courier: NewCourierRepo(pgpool),
		category: NewCategoryRepo(pgpool),
		product: NewProductRepo(pgpool),
		order: NewOrderRepo(pgpool),
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Customer() storage.CustomerRepoI {
	if s.customer == nil {
		s.customer = NewCustomerRepo(s.db)
	}
	return s.customer
}

func (s *Store) User() storage.UserRepoI {
	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}
	return s.user
}

func (s *Store) Courier() storage.CourierRepoI {
	if s.courier == nil {
		s.courier = NewCourierRepo(s.db)
	}
	return s.courier
}

func (s *Store) Category() storage.CategoryRepoI {
	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}
	return s.category
}

func (s *Store) Product() storage.ProductRepoI {
	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}
	return s.product
}

func (s *Store) Order() storage.OrderRepoI {
	if s.order == nil {
		s.order = NewOrderRepo(s.db)
	}
	return s.order
}