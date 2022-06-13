package repository

import (
	"context"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
	"github.com/gxrlxv/musique/auth_service/pkg/client/postgresql"
	"github.com/gxrlxv/musique/auth_service/pkg/logging"
)

type authRepository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) *authRepository {
	return &authRepository{
		client: client,
		logger: logger,
	}
}
func (ar *authRepository) Create(ctx context.Context, user domain.User) error {
	q := `insert into public.user (username, email, password, first_name, last_name, gender, country, city, phone) values ('asdasdasdsa',
  'gorlov@mail.ru', 'dasdasd', 'sergey', 'gorlov', 'male', 'russia', 'moscov', '890-63643') returning id`

	ar.client.QueryRow(ctx)
	return nil
}

func (ar *authRepository) FindByUsername(ctx context.Context, username string) (domain.User, error) {

	return domain.User{}, nil
}

func (ar *authRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {

	return domain.User{}, nil
}

func (ar *authRepository) FindByPhone(ctx context.Context, phone string) (domain.User, error) {

	return domain.User{}, nil
}
