package service

import (
	"backend/internal/pb"
	"backend/internal/utils"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ConnectorService struct {
}

func NewConnectorService() *ConnectorService {
	return &ConnectorService{}
}

func (c *ConnectorService) GetAllProjects(params *utils.PageParams) (*pb.AllProjectResponse, error) {
	conn, err := grpc.Dial("connector:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewJiraClient(conn)
	response, err := client.GetProjects(context.Background(), &pb.AllProjectsRequest{
		Limit:  int32(params.Limit),
		Page:   int32(params.Page),
		Search: params.Search,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *ConnectorService) DownloadProject(key string) (uint, error) {
	conn, err := grpc.Dial("connector:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	client := pb.NewJiraClient(conn)
	response, err := client.DownloadProject(context.Background(), &pb.DownloadProjectRequest{Key: key})
	return uint(response.Id), err
}
