package repository

import (
	"context"
	"fmt"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
	"github.com/gxrlxv/musique/auth_service/pkg/client/postgresql"
	"github.com/jackc/pgconn"
	"github.com/sirupsen/logrus"
)

type authRepository struct {
	client postgresql.Client
	log    *logrus.Logger
}

func NewAuthRepository(client postgresql.Client, log *logrus.Logger) *authRepository {
	return &authRepository{
		client: client,
		log:    log,
	}
}

func (ar *authRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	q := `
			INSERT INTO public.user 
				(username, email, password, first_name, last_name, gender, country, city, phone) 
		  	VALUES 
				($1, $2, $3, $4, $5, $6, $7, $8, $9) 
			RETURNING id`

	err := ar.client.QueryRow(ctx, q, user.Username, user.Email, user.PasswordHash, user.FirstName, user.LastName, user.Gender,
		user.Country, user.City, user.Phone).Scan(&user.ID)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf("SQL error: %s, Detail: %s, Where: %s, Code: %s, SQLstate: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			ar.log.Error(newErr)
			return &domain.User{}, newErr
		}
		return &domain.User{}, err
	}

	return user, nil
}

func (ar *authRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	q := `
		SELECT id, username, email, password, first_name, last_name, gender, country, city, phone, created_at, role 
		FROM public.user 
		WHERE username = $1
	`

	var u domain.User
	err := ar.client.QueryRow(ctx, q, username).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.FirstName,
		&u.LastName, &u.Gender, &u.Country, &u.City, &u.Phone, &u.CreatedAt, &u.Role)
	if err != nil {
		ar.log.Error(err)
		return &domain.User{}, err
	}

	return &u, nil
}

func (ar *authRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	q := `
		SELECT id, username, email, password, first_name, last_name, gender, country, city, phone, created_at, role 
		FROM public.user 
		WHERE email = $1
	`

	var u domain.User
	err := ar.client.QueryRow(ctx, q, email).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.FirstName,
		&u.LastName, &u.Gender, &u.Country, &u.City, &u.Phone, &u.CreatedAt, &u.Role)
	if err != nil {
		ar.log.Error(err)
		return &domain.User{}, err
	}

	return &u, nil
}

func (ar *authRepository) GetByPhone(ctx context.Context, phone string) (*domain.User, error) {
	q := `
		SELECT id, username, email, password, first_name, last_name, gender, country, city, phone, created_at, role 
		FROM public.user 
		WHERE phone = $1
	`

	var u domain.User
	err := ar.client.QueryRow(ctx, q, phone).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.FirstName,
		&u.LastName, &u.Gender, &u.Country, &u.City, &u.Phone, &u.CreatedAt, &u.Role)
	if err != nil {
		ar.log.Error(err)
		return &domain.User{}, err
	}

	return &u, nil
}

func (ar *authRepository) CreateSession(ctx context.Context, userID string) error {
	q := `
			INSERT INTO public.session
				(user_id)
			VALUES
				($1)`

	_, err := ar.client.Exec(ctx, q, userID)
	if err != nil {
		ar.log.Error(err)
		return err
	}

	return nil
}

func (ar *authRepository) UpdateSession(ctx context.Context, session *domain.Session) error {
	q := `
			UPDATE public.session
			SET refresh_token = $1, expires_at = $2
			WHERE user_id = $3`

	_, err := ar.client.Exec(ctx, q, session.RefreshToken, session.ExpiresAt, session.UserId)
	if err != nil {
		ar.log.Error(err)
		return err
	}

	return nil
}

func (ar *authRepository) GetIdByToken(ctx context.Context, refresh string) (string, error) {
	q := `
			SELECT user_id 
			FROM public.session 
			WHERE refresh_token = $1`

	var userId string
	err := ar.client.QueryRow(ctx, q, refresh).Scan(&userId)
	if err != nil {
		ar.log.Error(err)
		return "", err
	}

	return userId, nil
}
