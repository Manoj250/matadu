package grpc_handler

import (
	"context"
	pb "matadu/cache/internal/pb"
	"matadu/cache/internal/service"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CacheServiceServer struct {
	pb.UnimplementedCacheServiceServer
	service *service.CacheService
	logger  *zap.Logger
}

func NewCacheServiceServer(service *service.CacheService, zapLogger *zap.Logger) *CacheServiceServer {
	return &CacheServiceServer{
		service: service,
		logger:  zapLogger,
	}
}

func (c *CacheServiceServer) RegisterServer(ctx context.Context, req *pb.ServerDetails) (*pb.CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterServer not implemented")
}

func (c *CacheServiceServer) UnRegisterServer(ctx context.Context, req *pb.ServerDetails) (*pb.CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnRegisterServer not implemented")
}

func (c *CacheServiceServer) GetServerLoad(ctx context.Context, req *emptypb.Empty) (*pb.GetServerLoadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerLoad not implemented")
}

func (c *CacheServiceServer) GetUsersByChatId(ctx context.Context, req *pb.ChatDetails) (*pb.UsersByChatIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsersByChatId not implemented")
}

func (c *CacheServiceServer) RegisterUserToCache(ctx context.Context, req *pb.UserDetails) (*pb.CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUserToCache not implemented")
}

func (c *CacheServiceServer) UnRegisterUserFromCache(ctx context.Context, req *pb.UserDetails) (*pb.CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnRegisterUserFromCache not implemented")
}

func (c *CacheServiceServer) SetServerLoad(ctx context.Context, req *pb.SetServerLoadRequest) (*pb.CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetServerLoad not implemented")
}

func (c *CacheServiceServer) RegisterUserToChat(ctx context.Context, req *pb.ChatDetails) (*pb.CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUserToChat not implemented")
}

func (c *CacheServiceServer) UnRegisterUserFromChat(ctx context.Context, req *pb.ChatDetails) (*pb.CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnRegisterUserFromChat not implemented")
}
