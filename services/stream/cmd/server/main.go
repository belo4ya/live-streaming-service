package main

import (
	"context"
	"github.com/belo4ya/live-streaming-service/api/stream/v1"
	"github.com/belo4ya/live-streaming-service/third_party"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"io/fs"
	"log"
	"mime"
	"net"
	"net/http"
)

type Server struct {
	v1.UnimplementedStreamServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) ListStreams(_ context.Context, _ *emptypb.Empty) (*v1.ListStreamsResponse, error) {
	results := []*v1.ListStreamsResponse_Stream{
		{
			Id:       1,
			Name:     "Stream 1",
			Username: "Streamer 1",
		},
		{
			Id:       2,
			Name:     "Stream 2",
			Username: "Streamer 2",
		},
		{
			Id:       3,
			Name:     "Stream 3",
			Username: "Streamer 3",
		},
	}
	return &v1.ListStreamsResponse{Results: results}, nil
}

func RegisterDocHandler(mux *http.ServeMux, openAPISpec []byte) {
	mux.HandleFunc("/doc/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(openAPISpec)
	})

	mime.AddExtensionType(".svg", "image/svg+xml")
	subFs, _ := fs.Sub(third_party.SwaggerUI, "swagger-ui")
	mux.Handle("/doc/swagger-ui/", http.StripPrefix("/doc/swagger-ui", http.FileServer(http.FS(subFs))))
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	grpcSrv := grpc.NewServer()
	// Attach the Greeter service to the server
	v1.RegisterStreamServiceServer(grpcSrv, NewServer())
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(grpcSrv.Serve(lis))
	}()

	mux := http.NewServeMux()

	gw := runtime.NewServeMux()
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	// Register Greeter
	err = v1.RegisterStreamServiceHandler(context.Background(), gw, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	mux.Handle("/", gw)
	RegisterDocHandler(mux, v1.OpenAPISpec)

	gwSrv := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwSrv.ListenAndServe())
}
