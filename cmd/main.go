package main

import (
	"context"
	"link-shortener/internal/cfg"
	"link-shortener/internal/server"
	"link-shortener/internal/storage"
	cache "link-shortener/internal/storage/cache"
	pgdb "link-shortener/internal/storage/postgres"
	pb "link-shortener/pb"
	"net/http"

	"flag"
	"log"
	"net"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const path = "internal/cfg"

func main() {
	cfg.Init(path)

	var database string
	flag.StringVar(&database, "database", "in-memory", "select database: postgres | in-memory")
	flag.Parse()
	log.Printf("database selected: %s", database)
	var repo storage.StorageInstance

	switch database {
	case "in-memory":
		repo = cache.NewInMemory()
	case "postgres":
		p := new(pgdb.PostgresInstance)
		p.Setup()
		repo = pgdb.NewPostgresInstance(p)
	default:
		log.Fatal("non-existent repository specified")
	}
	go RunGatewayServer(cfg.ServConfig("HTTP"), repo)
	RunGrpcServer(cfg.ServConfig("gRPC"), repo)
}

func RunGrpcServer(servCfg cfg.Config, repo storage.StorageInstance) {
	s := grpc.NewServer()
	srv := server.NewServer(servCfg, repo)
	pb.RegisterLinkShortenerServer(s, srv)
	reflection.Register(s)

	listener, err := net.Listen("tcp", ":"+servCfg.Port)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("start gRPC gateaway server")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("cannot start gRPC server:%s", err)
	}

}

func RunGatewayServer(servCfg cfg.Config, repo storage.StorageInstance) {
	s := server.NewServer(servCfg, repo)
	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := pb.RegisterLinkShortenerHandlerServer(ctx, grpcMux, s)
	if err != nil {
		log.Fatal("cannot register handler server", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", servCfg.ServerAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("start HTTP gateaway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatalf("cannot start HTTP server:%s", err)
	}
}
