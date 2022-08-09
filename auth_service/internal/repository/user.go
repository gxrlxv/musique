package repository

import (
	"fmt"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
	"github.com/gxrlxv/musique/auth_service/pkg/client/postgresql"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUserNotFoundEmail    = status.Error(codes.NotFound, "user with given email not found")
	ErrUserNotFoundUsername = status.Error(codes.NotFound, "user with given username not found")
	ErrUserNotFoundPhone    = status.Error(codes.NotFound, "user with given phone not found")
)

type userRepository struct {
	client postgresql.Client
	log    *logrus.Logger
}

func NewUserRepository(client postgresql.Client, log *logrus.Logger) *userRepository {
	return &userRepository{
		client: client,
		log:    log,
	}
}

func (ur *userRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	q := `
			INSERT INTO public.user 
				(username, email, password, first_name, last_name, gender, country, city, phone) 
		  	VALUES 
				($1, $2, $3, $4, $5, $6, $7, $8, $9) 
			RETURNING id`

	err := ur.client.QueryRow(ctx, q, user.Username, user.Email, user.PasswordHash, user.FirstName, user.LastName, user.Gender,
		user.Country, user.City, user.Phone).Scan(&user.ID)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf("SQL error: %s, Detail: %s, Where: %s, Code: %s, SQLstate: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			ur.log.Error(newErr)
			return domain.User{}, newErr
		}
		return domain.User{}, err
	}

	return user, err
}

func (ur *userRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	q := `
		SELECT id, username, email, password, first_name, last_name, gender, country, city, phone, created_at, role 
		FROM public.user 
		WHERE username = $1
	`

	var u domain.User
	err := ur.client.QueryRow(ctx, q, username).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.FirstName,
		&u.LastName, &u.Gender, &u.Country, &u.City, &u.Phone, &u.CreatedAt, &u.Role)
	if err != nil {
		ur.log.Error(err)
		if err == pgx.ErrNoRows {
			return domain.User{}, ErrUserNotFoundUsername
		}
		return domain.User{}, err
	}

	return u, err
}

func (ur *userRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	q := `
		SELECT id, username, email, password, first_name, last_name, gender, country, city, phone, created_at, role 
		FROM public.user 
		WHERE email = $1
	`

	var u domain.User
	err := ur.client.QueryRow(ctx, q, email).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.FirstName,
		&u.LastName, &u.Gender, &u.Country, &u.City, &u.Phone, &u.CreatedAt, &u.Role)
	if err != nil {
		ur.log.Error(err)
		if err == pgx.ErrNoRows {
			return domain.User{}, ErrUserNotFoundEmail
		}
		return domain.User{}, err
	}

	return u, err
}

func (ur *userRepository) GetByPhone(ctx context.Context, phone string) (domain.User, error) {
	q := `
		SELECT id, username, email, password, first_name, last_name, gender, country, city, phone, created_at, role 
		FROM public.user 
		WHERE phone = $1
	`

	var u domain.User
	err := ur.client.QueryRow(ctx, q, phone).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.FirstName,
		&u.LastName, &u.Gender, &u.Country, &u.City, &u.Phone, &u.CreatedAt, &u.Role)
	if err != nil {
		ur.log.Error(err)
		if err == pgx.ErrNoRows {
			return domain.User{}, ErrUserNotFoundPhone
		}
		return domain.User{}, err
	}

	return u, err
}
