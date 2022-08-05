package app

import (
	"context"
	"github.com/gxrlxv/musique/artist_service/internal/config"
	"github.com/gxrlxv/musique/artist_service/internal/repository"
	"github.com/gxrlxv/musique/artist_service/internal/server"
	"github.com/gxrlxv/musique/artist_service/internal/service"
	"github.com/gxrlxv/musique/artist_service/internal/usecase"
	"github.com/gxrlxv/musique/artist_service/pkg/client/postgresql"
	"github.com/gxrlxv/musique/artist_service/pkg/logging"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
)

func Run() {
	log := logging.GetLogger()

	cfg := config.GetConfig()

	postgreSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		log.Fatalf("%v", err)
	}

	albumRepo := repository.NewAlbumRepository(postgreSQLClient, log)

	trackRepo := repository.NewTrackRepository(postgreSQLClient, log)

	genreRepo := repository.NewGenreRepository(postgreSQLClient, log)

	artistRepo := repository.NewArtistRepository(postgreSQLClient, log)

	albumUseCase := usecase.NewAlbumUseCase(albumRepo, trackRepo, genreRepo, artistRepo, log)

	artistService := service.NewArtistService(albumUseCase, log)

	grpcServer := server.NewGRPCServer(artistService, log)

	listen, err := net.Listen("tcp", cfg.Server.Grpc.Addr)
	if err != nil {
		log.Panic(err)
	}

	mux := server.NewHTTPServer(cfg.Server.Grpc.Addr, log)

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() (err error) {
		return grpcServer.Serve(listen)
	})
	g.Go(func() (err error) {
		return http.ListenAndServe(cfg.Server.Http.Addr, mux)
	})

	err = g.Wait()
	if err != nil {
		log.Panic(err)
	}
}
