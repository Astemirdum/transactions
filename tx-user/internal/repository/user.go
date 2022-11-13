package repository

import (
	"context"

	models "github.com/Astemirdum/transactions/tx-user/models/v1"
	"github.com/jmoiron/sqlx"
)

func (u *UserDB) CreateUser(ctx context.Context, user models.User, tx *sqlx.Tx) (int, error) {
	query, args, err := psql.Insert(userTable).
		Columns("email", "hash_password").
		Values(user.Email, user.Password).
		Suffix("RETURNING id").ToSql()
	if err != nil {
		return 0, err
	}
	var id int
	row := tx.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UserDB) GetUser(ctx context.Context, email, password string) (models.User, error) {
	query, args, err := psql.
		Select("id", "email", "hash_password").
		From(userTable).
		Where("email=$1", email).
		Where("hash_password=$2", password).ToSql()
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	if err := u.db.GetContext(ctx, &user, query, args...); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *UserDB) ExistsEmail(ctx context.Context, email string) (bool, error) {
	query, args, err := psql.
		Select("count(*)").
		From(userTable).
		Where("email=$1", email).ToSql()
	if err != nil {
		return false, err
	}
	var count int
	if err := u.db.GetContext(ctx, &count, query, args...); err != nil {
		return false, err
	}
	return count > 0, nil
}
