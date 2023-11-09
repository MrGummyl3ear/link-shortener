package server

import (
	"context"
	"fmt"
	"link-shortener/internal/cfg"
	"link-shortener/internal/storage"
	pb "link-shortener/pb"
	"log"
)

// GRPCServer...
type GRPCServer struct {
	db  storage.StorageInstance
	Cfg cfg.Config
	pb.UnimplementedLinkShortenerServer
}

func NewServer(cfg cfg.Config, repo storage.StorageInstance) *GRPCServer {
	return &GRPCServer{db: repo, Cfg: cfg}
}

// Put...
func (s *GRPCServer) PutUrl(ctx context.Context, req *pb.Link) (*pb.Link, error) {
	//newUrl := Gen_ShortUrl(s.Cfg.ServerAddress, utils.Hash_func(req.Url, s.Cfg.LinkLength))
	newUrl, err := s.db.Save(req.Url, s.Cfg.LinkLength)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newUrl)
	return &pb.Link{Url: newUrl}, nil
}

// Get...
func (s *GRPCServer) GetUrl(ctx context.Context, req *pb.Link) (*pb.Link, error) {
	newUrl, err := s.db.Get(req.Url)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.Link{Url: newUrl}, nil
}

/*
func Gen_ShortUrl(baseUrl string, url string) string {
	return baseUrl + "/" + url
}
*/
