package server

import (
	"context"
	"link-shortener/internal/cfg"
	"link-shortener/internal/storage"
	"link-shortener/internal/utils"
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
	var shortUrl, copyUrl string
	copyUrl = req.Url
	for {
		shortUrl = utils.Hash(copyUrl, s.Cfg.LinkLength)
		if s.db.Unique(shortUrl, copyUrl) {
			break
		} else {
			copyUrl += shortUrl
		}
	}
	err := s.db.Save(copyUrl, shortUrl)
	if err != nil {
		log.Printf("error with saving the url:%s", err)
		return &pb.Link{Url: ""}, err
	}
	return &pb.Link{Url: shortUrl}, nil
}

// Get...
func (s *GRPCServer) GetUrl(ctx context.Context, req *pb.Link) (*pb.Link, error) {
	newUrl, err := s.db.Get(req.Url)
	if err != nil {
		log.Println(err)
		return &pb.Link{Url: ""}, err
	}
	return &pb.Link{Url: newUrl}, nil
}
