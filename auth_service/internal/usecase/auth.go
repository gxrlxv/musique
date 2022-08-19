package usecase

import (
	"context"
	v1 "github.com/gxrlxv/musique/auth_service/api/auth/v1"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
	"github.com/gxrlxv/musique/auth_service/internal/repository"
	"github.com/gxrlxv/musique/auth_service/pkg/auth"
	"github.com/gxrlxv/musique/auth_service/pkg/hash"
	"github.com/sirupsen/logrus"
	"time"
)

var artistRole = "artist"

type authUseCase struct {
	userRepo         repository.UserRepository
	sessionRepo      repository.SessionRepository
	playlistRepo     repository.PlaylistRepository
	subscriptionRepo repository.SubscriptionRepository
	log              *logrus.Logger
	hasher           hash.PasswordHasher
	tokenManager     auth.Manager
	accessTokenTTL   time.Duration
	refreshTokenTTL  time.Duration
}

func NewAuthUseCase(userRepo repository.UserRepository, sessionRepo repository.SessionRepository,
	playlistRepo repository.PlaylistRepository, subscriptionRepo repository.SubscriptionRepository,
	log *logrus.Logger, hasher hash.PasswordHasher, tokenManager auth.Manager,
	accessTokenTTL time.Duration, refreshTokenTTL time.Duration) *authUseCase {
	return &authUseCase{
		userRepo:         userRepo,
		sessionRepo:      sessionRepo,
		playlistRepo:     playlistRepo,
		subscriptionRepo: subscriptionRepo,
		log:              log,
		hasher:           hasher,
		tokenManager:     tokenManager,
		accessTokenTTL:   accessTokenTTL,
		refreshTokenTTL:  refreshTokenTTL,
	}
}

func (a *authUseCase) SignUp(ctx context.Context, dto domain.CreateUserDTO) (domain.User, error) {
	a.log.Info("signUp use case")
	if dto.Password != dto.RepeatPassword {
		return domain.User{}, ErrPasswordDontMatch
	}

	if u, _ := a.userRepo.GetByEmail(ctx, dto.Email); u.Email == dto.Email {
		a.log.Error(ErrUserAlreadyExistEmail)
		return domain.User{}, ErrUserAlreadyExistEmail
	}

	if u, _ := a.userRepo.GetByUsername(ctx, dto.Username); u.Username == dto.Username {
		a.log.Error(ErrUserAlreadyExistUsername)
		return domain.User{}, ErrUserAlreadyExistUsername
	}

	if u, _ := a.userRepo.GetByPhone(ctx, dto.Phone); u.Phone == dto.Phone {
		a.log.Error(ErrUserAlreadyExistPhone)
		return domain.User{}, ErrUserAlreadyExistPhone
	}
	a.log.Info("hash password")
	passwordHash, err := a.hasher.Hash(dto.Password)
	if err != nil {
		a.log.Error(err)
		return domain.User{}, internalErr(err)
	}

	model := domain.NewUser(dto, passwordHash)

	user, err := a.userRepo.Create(ctx, *model)
	if err != nil {
		a.log.Error(err)
		return domain.User{}, internalErr(err)
	}

	err = a.sessionRepo.CreateSession(ctx, user.ID)
	if err != nil {
		a.log.Error(err)
		return domain.User{}, internalErr(err)
	}

	return user, err
}

func (a *authUseCase) SignIn(ctx context.Context, email, password string) (domain.User, error) {
	a.log.Info("signIn use case")
	user, err := a.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	passwordHash, err := a.hasher.Hash(password)
	if err != nil {
		a.log.Error(err)
		return domain.User{}, internalErr(err)
	}

	if user.PasswordHash != passwordHash {
		a.log.Error(ErrPasswordInvalid)
		return domain.User{}, ErrPasswordInvalid
	}

	return user, err
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

	if err := a.sessionRepo.UpdateSession(ctx, *session); err != nil {
		a.log.Error(err)
		return &v1.Tokens{}, internalErr(err)
	}

	return &v1.Tokens{AccessToken: access, RefreshToken: refresh}, err
}

func (a *authUseCase) GetIdFromRefresh(ctx context.Context, refresh string) (string, error) {
	a.log.Info("GetIdFromRefresh use case")

	userId, err := a.sessionRepo.GetIdByToken(ctx, refresh)
	if err != nil {
		a.log.Error(ErrTokenInvalid)
		return "", ErrTokenInvalid
	}

	return userId, err
}

func (a *authUseCase) Identify(ctx context.Context, access string) (string, string, error) {
	a.log.Info("Identify artist use case")

	claims, err := a.tokenManager.Parse(access)
	if err != nil {
		a.log.Error(ErrTokenInvalid)
		return "", "", ErrTokenInvalid
	}

	return claims["sub"].(string), claims["role"].(string), err
}
