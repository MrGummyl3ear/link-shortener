package main

import (
	"link-shortener/internal/cfg"
	"link-shortener/internal/server"
	"link-shortener/internal/storage"
	cache "link-shortener/internal/storage/cache"
	pgdb "link-shortener/internal/storage/postgres"
	pb "link-shortener/pb"

	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const path = "internal/cfg"

func main() {
	cfg.Init(path)
	servCfg := cfg.ServConfig()

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
	s := grpc.NewServer()
	srv := server.NewServer(servCfg, repo)
	pb.RegisterLinkShortenerServer(s, srv)
	reflection.Register(s)

	l, err := net.Listen("tcp", ":"+servCfg.Port)

	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
