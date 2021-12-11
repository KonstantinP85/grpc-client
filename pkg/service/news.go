package service

import (
	"context"
	grpc_client "github.com/KonstantinP85/grpc-client"
	"github.com/KonstantinP85/grpc-client/pkg/api"
	"log"
	"time"
)

type NewsService struct {
}

func NewNewsService() *NewsService {
	return &NewsService{}
}

func (s *NewsService) UploadNews() (int32, error) {

	c := runGrpcClient()
	data, err := c.UploadNews(context.Background(), &api.UploadNewsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	return data.Count, nil
}

func (s *NewsService) GetNews(id int) (grpc_client.News, error) {

	c := runGrpcClient()
	data, err := c.GetNews(context.Background(), &api.GetNewsRequest{
		Id: int32(id),
	})
	if err != nil {
		log.Fatal(err)
	}
	tm := time.Unix(data.GetNews().Datetime, 0)

	news := grpc_client.News{
		Id:       data.GetNews().Id,
		Category: data.GetNews().Category,
		Datetime: tm.Format("2006-01-02 15:04:05"),
		Headline: data.GetNews().Headline,
		Image:    data.GetNews().Image,
		Related:  data.GetNews().Related,
		Resource: data.GetNews().Resource,
		Summary:  data.GetNews().Summary,
		Url:      data.GetNews().Url,
	}

	return news, nil
}
