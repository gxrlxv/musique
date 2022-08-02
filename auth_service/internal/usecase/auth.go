package usecase

import (
	"context"
	v1 "github.com/gxrlxv/musique/auth_service/api/auth/v1"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
	"github.com/gxrlxv/musique/auth_service/pkg/auth"
	"github.com/gxrlxv/musique/auth_service/pkg/hash"
	"github.com/sirupsen/logrus"
	"time"
)

type AuthRepository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	GetByPhone(ctx context.Context, phone string) (*domain.User, error)
	UpdateSession(ctx context.Context, session *domain.Session) error
	CreateSession(ctx context.Context, userID string) error
	GetIdByToken(ctx context.Context, refresh string) (string, error)
}

type authUseCase struct {
	authRepo        AuthRepository
	playlistRepo    PlaylistRepository
	hasher          hash.PasswordHasher
	tokenManager    auth.Manager
	log             *logrus.Logger
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewAuthUseCase(authRepo AuthRepository, playlistRepo PlaylistRepository, hasher hash.PasswordHasher, tokenManager auth.Manager, log *logrus.Logger,
	accessTokenTTL time.Duration, refreshTokenTTL time.Duration) *authUseCase {
	return &authUseCase{
		authRepo:        authRepo,
		playlistRepo:    playlistRepo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		log:             log,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

func (a *authUseCase) SignUp(ctx context.Context, dto domain.CreateUserDTO) (*domain.User, error) {
	a.log.Info("signUp use case")
	if dto.Password != dto.RepeatPassword {
		return &domain.User{}, ErrPasswordDontMatch
	}

	if u, _ := a.authRepo.GetByEmail(ctx, dto.Email); u.Email == dto.Email {
		a.log.Error(ErrUserAlreadyExistEmail)
		return &domain.User{}, ErrUserAlreadyExistEmail
	}

	if u, _ := a.authRepo.GetByUsername(ctx, dto.Username); u.Username == dto.Username {
		a.log.Error(ErrUserAlreadyExistUsername)
		return &domain.User{}, ErrUserAlreadyExistUsername
	}

	if u, _ := a.authRepo.GetByPhone(ctx, dto.Phone); u.Phone == dto.Phone {
		a.log.Error(ErrUserAlreadyExistPhone)
		return &domain.User{}, ErrUserAlreadyExistPhone
	}
	a.log.Info("hash password")
	passwordHash, err := a.hasher.Hash(dto.Password)
	if err != nil {
		a.log.Error(err)
		return &domain.User{}, internalErr(err)
	}

	model := domain.NewUser(dto, passwordHash)

	user, err := a.authRepo.Create(ctx, model)
	if err != nil {
		a.log.Error(err)
		return &domain.User{}, internalErr(err)
	}

	err = a.authRepo.CreateSession(ctx, user.ID)
	if err != nil {
		a.log.Error(err)
		return &domain.User{}, internalErr(err)
	}

	return user, nil
}

func (a *authUseCase) SignIn(ctx context.Context, email, password string) (*domain.User, error) {
	a.log.Info("signIn use case")
	user, err := a.authRepo.GetByEmail(ctx, email)
	if err != nil {
		a.log.Error(ErrUserNotFoundEmail)
		return &domain.User{}, ErrUserNotFoundEmail
	}

	passwordHash, err := a.hasher.Hash(password)
	if err != nil {
		a.log.Error(err)
		return &domain.User{}, internalErr(err)
	}

	if user.PasswordHash != passwordHash {
		a.log.Error(ErrPasswordInvalid)
		return &domain.User{}, ErrPasswordInvalid
	}

	return user, nil
}

func (a *authUseCase) NewTokens(ctx context.Context, userId, role string) (*v1.Tokens, error) {
	a.log.Info("NewTokens use case")
	refresh, err := a.tokenManager.NewRefreshToken()
	if err != nil {
		a.log.Error(err)
		return &v1.Tokens{}, internalErr(err)
	}

	access, err := a.tokenManager.NewJWT(userId, role, a.accessTokenTTL)
	if err != nil {
		a.log.Error(err)
		return &v1.Tokens{}, internalErr(err)
	}

	session := domain.NewSession(userId, refresh, a.refreshTokenTTL)

	if err := a.authRepo.UpdateSession(ctx, session); err != nil {
		a.log.Error(err)
		return &v1.Tokens{}, internalErr(err)
	}

	return &v1.Tokens{AccessToken: access, RefreshToken: refresh}, nil
}

func (a *authUseCase) GetIdFromRefresh(ctx context.Context, refresh string) (string, error) {
	a.log.Info("GetIdFromRefresh use case")
	userId, err := a.authRepo.GetIdByToken(ctx, refresh)
	if err != nil {
		a.log.Error(ErrTokenInvalid)
		return "", ErrTokenInvalid
	}

	return userId, nil
}

func (a *authUseCase) Identify(ctx context.Context, access string) (string, error) {
	a.log.Info("Identify use case")
	userID, err := a.tokenManager.Parse(access)
	if err != nil {
		a.log.Error(ErrTokenInvalid)
		return "", ErrTokenInvalid
	}

	return userID, nil
}
