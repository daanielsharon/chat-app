package repository

import (
	"context"
	"database/sql"
	"server/model/domain"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type UserRepositoryImpl struct {
	db DBTX
}

func NewRepository(db DBTX) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *domain.User) (*domain.User, error)  {
	var lastInsertId int
	query := "INSERT INTO users(username, email, password) VALUES($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&lastInsertId)

	if err != nil {
		return &domain.User{}, err
	}

	user.ID = int64(lastInsertId)
	return user, nil
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	u := domain.User{}

	query := `SELECT id, username, email, password FROM users WHERE email = $1`
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Username, &u.Email, &u.Password)

	if err != nil {
		return &domain.User{}, err
	}
	
	return &u, nil
}