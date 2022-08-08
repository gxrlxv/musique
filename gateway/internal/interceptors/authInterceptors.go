package interceptors

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/gxrlxv/musique/auth_service/api/auth/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	authorizationHeader = runtime.MetadataPrefix + "authorization"
	userIDHeader        = "user-id"
)

type AuthInterceptor struct {
	client v1.AuthClient
}

func NewAuthInterceptor(addr string, log *logrus.Logger) (*AuthInterceptor, func(), error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		log.Info("closing the auth connection")
		if err := conn.Close(); err != nil {
			log.Errorf("error while closing auth connection: %v", err)
		}
	}

	return &AuthInterceptor{v1.NewAuthClient(conn)}, cleanup, nil
}

func (i *AuthInterceptor) identityArtist(ctx context.Context) (context.Context, error) {
	val := metautils.ExtractOutgoing(ctx).Get(authorizationHeader)
	reply, err := i.client.IdentifyArtist(ctx, &v1.IdentifyArtistRequest{AccessToken: strings.TrimPrefix(val, "Bearer ")})
	if err != nil {
		return nil, err
	}

	return metadata.AppendToOutgoingContext(ctx, userIDHeader, reply.Role), nil
}

func (i *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		ctx, err := i.identityArtist(ctx)
		if err != nil {
			return err
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func (i *AuthInterceptor) Stream() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		ctx, err := i.identityArtist(ctx)
		if err != nil {
			return nil, err
		}
		return streamer(ctx, desc, cc, method, opts...)
	}
}
