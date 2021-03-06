package user

import (
	"context"
	"dojo_go/config/database"
	"dojo_go/model"
	"dojo_go/repository"
	"time"
)

// NewSQLPostRepo returns implement of user repository interface
func NewPostgresUserRepo(Conn *database.Data) repository.UserRepository {
	return &sqlUserRepo{
		Conn: Conn,
	}
}

type sqlUserRepo struct {
	Conn *database.Data
}

// GetAll returns all users.
func (ur *sqlUserRepo) GetAllUsers(ctx context.Context) ([]model.User, error) {
	rows, err := ur.Conn.DB.QueryContext(ctx, selectAllUser)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var userRow model.User
		_ = rows.Scan(&userRow.ID, &userRow.Name, &userRow.Surname, &userRow.Username, &userRow.CreatedAt, &userRow.UpdatedAt)
		createdAtUnix := userRow.CreatedAt.Unix()
		userRow.CreatedAtInt = uint64(createdAtUnix)
		updatedAtUnix := userRow.UpdatedAt.Unix()
		userRow.UpdatedAtInt = uint64(updatedAtUnix)
		users = append(users, userRow)
	}

	return users, nil
}

// GetOne returns one user by id.
func (ur *sqlUserRepo) GetOne(ctx context.Context, id uint) (model.User, error) {
	row := ur.Conn.DB.QueryRowContext(ctx, selectUserById, id)

	var userScan model.User
	err := row.Scan(&userScan.ID, &userScan.Name, &userScan.Username, &userScan.Username, &userScan.CreatedAt, &userScan.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}

	createdAtUnix := userScan.CreatedAt.Unix()
	userScan.CreatedAtInt = uint64(createdAtUnix)

	updatedAtUnix := userScan.UpdatedAt.Unix()
	userScan.UpdatedAtInt = uint64(updatedAtUnix)

	return userScan, nil
}

// Create adds a new user.
func (ur *sqlUserRepo) Create(ctx context.Context, user *model.User) error {
	now := time.Now().Truncate(time.Second).Truncate(time.Millisecond).Truncate(time.Microsecond)

	stmt, err := ur.Conn.DB.PrepareContext(ctx, insertUser)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, user.Name, user.Username, user.Username, user.PasswordHash, now, now)
	err = row.Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *sqlUserRepo) GetByUsername(ctx context.Context, username string) (model.User, error) {
	panic("implement me")
}

func (ur *sqlUserRepo) Update(ctx context.Context, id uint, user model.User) error {
	panic("implement me")
}

func (ur *sqlUserRepo) Delete(ctx context.Context, id uint) error {
	panic("implement me")
}
