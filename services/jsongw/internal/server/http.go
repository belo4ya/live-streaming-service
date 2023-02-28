package server

import (
	"context"
	v1 "github.com/belo4ya/live-streaming-service/api/jsongw/v1"
	"github.com/belo4ya/live-streaming-service/services/jsongw/internal/conf"
	"github.com/belo4ya/live-streaming-service/third_party"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io/fs"
	"mime"
	"net/http"
)

func NewHTTPServer(c *conf.HTTP, conn *grpc.ClientConn, logger log.Logger) *khttp.Server {
	mux := http.NewServeMux()
	gw := runtime.NewServeMux()
	err := v1.RegisterStreamServiceHandler(context.Background(), gw, conn)
	if err != nil {
		panic(err)
	}
	mux.Handle("/", gw)
	OpenAPIHandler(mux)

	var opts = []khttp.ServerOption{
		khttp.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
	}
	opts = append(opts, khttp.Address(c.Addr))
	srv := khttp.NewServer(opts...)
	srv.HandlePrefix("/", CORSHandler(mux))
	return srv
}

func NewGRPCConn(c *conf.GRPC) (*grpc.ClientConn, error) {
	conn, err := grpc.DialContext(
		context.Background(),
		c.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func OpenAPIHandler(mux *http.ServeMux) http.Handler {
	mux.HandleFunc("/doc/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(v1.OpenAPISpec)
	})
	mime.AddExtensionType(".svg", "image/svg+xml")
	subFs, _ := fs.Sub(third_party.SwaggerUI, "swagger-ui")
	mux.Handle("/doc/swagger-ui/", http.StripPrefix("/doc/swagger-ui", http.FileServer(http.FS(subFs))))
	return mux
}

func CORSHandler(h http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace,
		},
	})
	return c.Handler(h)
}
