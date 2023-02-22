package server

import (
	"context"
	v1 "github.com/belo4ya/live-streaming-service/api/vod/v1"
	"github.com/belo4ya/live-streaming-service/third_party"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io/fs"
	"log"
	"mime"
	"net/http"
)

func NewHTTPServer(addr string, grpcAddr string) *http.Server {
	gw := runtime.NewServeMux()
	conn, err := grpc.DialContext(
		context.Background(),
		grpcAddr,
		//grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	err = v1.RegisterVODServiceHandler(context.Background(), gw, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", gw)
	registerDocHandler(mux, v1.OpenAPISpec)

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}

func registerDocHandler(mux *http.ServeMux, openAPISpec []byte) {
	mux.HandleFunc("/doc/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(openAPISpec)
	})

	mime.AddExtensionType(".svg", "image/svg+xml")
	subFs, _ := fs.Sub(third_party.SwaggerUI, "swagger-ui")
	mux.Handle("/doc/swagger-ui/", http.StripPrefix("/doc/swagger-ui", http.FileServer(http.FS(subFs))))
}
