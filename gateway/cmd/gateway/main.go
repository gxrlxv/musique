package main

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/gxrlxv/musique/gateway/internal/config"
	"github.com/gxrlxv/musique/gateway/internal/interceptors"
	"github.com/gxrlxv/musique/gateway/internal/registers"
	"github.com/gxrlxv/musique/gateway/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"net/http"
)

func main() {
	log := logging.GetLogger()

	cfg := config.GetConfig()

	authInterceptors, cleanup, err := interceptors.NewAuthInterceptor(cfg.Auth.Addr, log)
	if err != nil {
		log.Error(err)
	}

	serverMux, cleanup2, err := registers.RegisterAll(cfg.Auth.Addr, cfg.Artist.Addr, authInterceptors)
	if err != nil {
		log.Error(err)
	}
	cleanupA := func() {
		cleanup()
		cleanup2()
	}
	defer cleanupA()
	ochttpHandler, err := customHandler(serverMux, log)
	if err != nil {
		log.Error(err)
	}

	server := newGatewayServer(cfg.Rest.Addr, ochttpHandler)

	err = server.ListenAndServe()
	if err != nil {
		log.Error(err)
	}

}
func newGatewayServer(addr string, oc http.Handler) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: oc,
	}
}

func customHandler(mux *runtime.ServeMux, log *logrus.Logger) (http.Handler, error) {
	if err := mux.HandlePath(
		"GET",
		"/health",
		func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			if _, err := w.Write([]byte("Gateway v1 serving")); err != nil {
				log.Errorf("Health check error: %v", err)
			}
		},
	); err != nil {
		return nil, err
	}

	return wsproxy.WebsocketProxy(mux), nil
}
