package service

import (
	grpc_client "github.com/KonstantinP85/grpc-client"
	"github.com/KonstantinP85/grpc-client/pkg/api"
	"google.golang.org/grpc"
	"log"
)

type News interface {
	GetNews(id int) (grpc_client.News, error)
	UploadNews() (int32, error)
}

type Service struct {
	News
}

func NewService() *Service {
	return &Service{
		News: NewNewsService(),
	}
}

func runGrpcClient() api.NewsServiceClient {

	conn, err := grpc.Dial(":8010", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewNewsServiceClient(conn)

	return c
}
