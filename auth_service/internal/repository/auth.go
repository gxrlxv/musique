package repository

import (
	"context"
	"fmt"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
	"github.com/gxrlxv/musique/auth_service/pkg/client/postgresql"
	"github.com/gxrlxv/musique/auth_service/pkg/logging"
	"github.com/jackc/pgconn"
	"strings"
)

type authRepository struct {
	client postgresql.Client
	log    *logging.Logger
}

func NewRepository(client postgresql.Client, log *logging.Logger) *authRepository {
	return &authRepository{
		client: client,
		log:    log,
	}
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (ar *authRepository) Create(ctx context.Context, user *domain.User) error {
	q := `
			insert into public.user 
				(username, email, password, first_name, last_name, gender, country, city, phone) 
		  	values 
				($1, $2, $3, $4, $5, $6, $7, $8, $9) 
			returning id`
	ar.log.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	err := ar.client.QueryRow(ctx, q, user.Username, user.Email, user.PasswordHash, user.FirstName, user.LastName, user.Gender,
		user.Country, user.City, user.Phone).Scan(&user.ID)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf("SQL error: %s, Detail: %s, Where: %s, Code: %s, SQLstate: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			ar.log.Error(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (ar *authRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	q := `
		SELECT id, username, email, password, first_name, last_name, gender, country, city, phone, created_at, role FROM public.user WHERE username = $1
	`
	ar.log.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var u domain.User
	err := ar.client.QueryRow(ctx, q, username).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.FirstName,
		&u.LastName, &u.Gender, &u.Country, &u.City, &u.Phone, &u.CreatedAt, &u.Role)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (ar *authRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	q := `
		SELECT id, username, email, password, first_name, last_name, gender, country, city, phone, created_at, role FROM public.user WHERE email = $1
	`
	ar.log.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var u domain.User
	err := ar.client.QueryRow(ctx, q, email).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.FirstName,
		&u.LastName, &u.Gender, &u.Country, &u.City, &u.Phone, &u.CreatedAt, &u.Role)
	if err != nil {
		ar.log.Info(err)
		return domain.User{}, err
	}

	return u, nil
}

func (ar *authRepository) GetByPhone(ctx context.Context, phone string) (domain.User, error) {
	q := `
		SELECT id, username, email, password, first_name, last_name, gender, country, city, phone, created_at, role FROM public.user WHERE phone = $1
	`
	ar.log.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var u domain.User
	err := ar.client.QueryRow(ctx, q, phone).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.FirstName,
		&u.LastName, &u.Gender, &u.Country, &u.City, &u.Phone, &u.CreatedAt, &u.Role)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (ar *authRepository) SetSession(ctx context.Context, session *domain.Session) error {
	q := `
			insert into public.session
				(user_id, refresh_token, expires_at)
			values
				($1, $2, $3)`
	ar.log.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	_, err := ar.client.Exec(ctx, q, session.UserId, session.RefreshToken, session.ExpiresAt)
	if err != nil {
		return err
	}

	return nil
}
