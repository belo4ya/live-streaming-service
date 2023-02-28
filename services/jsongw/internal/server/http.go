package server

import (
	"context"
	v1 "github.com/belo4ya/live-streaming-service/api/jsongw/v1"
	"github.com/belo4ya/live-streaming-service/services/jsongw/internal/conf"
	"github.com/belo4ya/live-streaming-service/third_party"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
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
	mux.Handle("/doc/", http.StripPrefix("/doc", OpenAPIHandler()))

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

func NewGRPCConn(c *conf.GRPC) *grpc.ClientConn {
	conn, err := kgrpc.DialInsecure(
		context.Background(),
		kgrpc.WithEndpoint(c.Addr),
		kgrpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return conn
}

func OpenAPIHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(v1.OpenAPISpec)
	})
	mime.AddExtensionType(".svg", "image/svg+xml")
	subFs, _ := fs.Sub(third_party.SwaggerUI, "swagger-ui")
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", http.FileServer(http.FS(subFs))))
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
