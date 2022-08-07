package registers

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	artistv1 "github.com/gxrlxv/musique/artist_service/api/artist/v1"
	authv1 "github.com/gxrlxv/musique/auth_service/api/auth/v1"
	"github.com/gxrlxv/musique/gateway/internal/interceptors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterAll(authAddr string, artistAddr string, authInt *interceptors.AuthInterceptor) (*runtime.ServeMux, func(), error) {
	ctx, cancel := context.WithCancel(context.Background())

	cleanup := func() {
		cancel()
	}

	mux := runtime.NewServeMux()

	if err := registerAuth(ctx, mux, authAddr); err != nil {
		return nil, cleanup, err
	}
	if err := registerArtist(ctx, mux, artistAddr, authInt); err != nil {
		return nil, cleanup, err
	}

	return mux, cleanup, nil
}

func registerAuth(ctx context.Context, mux *runtime.ServeMux, addr string) error {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	return authv1.RegisterAuthHandlerFromEndpoint(ctx, mux, addr, opts)
}

func registerArtist(ctx context.Context, mux *runtime.ServeMux, addr string, authInt *interceptors.AuthInterceptor) error {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(authInt.Unary()),
		grpc.WithStreamInterceptor(authInt.Stream()),
	}
	return artistv1.RegisterArtistHandlerFromEndpoint(ctx, mux, addr, opts)
}
