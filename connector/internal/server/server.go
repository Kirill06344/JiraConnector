package server

import (
	"context"
	"github.com/stewie/internal/connector"
	"github.com/stewie/internal/pb"
	"github.com/stewie/internal/utils"
	_ "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedJiraServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) GetProjects(ctx context.Context, request *pb.AllProjectsRequest) (*pb.AllProjectResponse, error) {
	params, err := utils.GetQueryParams(request)

	if err != nil {
		err = status.Error(codes.InvalidArgument, err.Error())
		return nil, err
	}
	result, err := connector.GetProjects(params)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}
	return result, nil
}

func (s *server) DownloadProject(ctx context.Context, request *pb.DownloadProjectRequest) (*pb.DownloadProjectResponse, error) {
	id, err := connector.DownloadProject(request.Key)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}

	return &pb.DownloadProjectResponse{Id: uint32(id)}, nil
}
