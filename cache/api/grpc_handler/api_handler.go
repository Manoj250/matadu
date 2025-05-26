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
	err := c.service.RegisterServer(ctx, req)
	if err != nil {
		c.logger.Error("failed to add server", zap.Error(err))
		return &pb.CommonResponse{
			Message: "Failed to register server to cache ",
			Code:    int32(codes.Internal),
			Success: false,
		}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pb.CommonResponse{
		Message: "server registered to cache",
		Code:    int32(codes.OK),
		Success: true,
	}, nil
}

func (c *CacheServiceServer) UnRegisterServer(ctx context.Context, req *pb.ServerDetails) (*pb.CommonResponse, error) {
	err := c.service.UnRegisterServer(ctx, req)
	if err != nil {
		c.logger.Error("failed to yeet the server", zap.Error(err))
		return &pb.CommonResponse{
			Message: "Failed to yeet the server from cache ",
			Code:    int32(codes.Internal),
			Success: false,
		}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pb.CommonResponse{
		Message: "server yeeted from cache",
		Code:    int32(codes.OK),
		Success: true,
	}, nil
}

func (c *CacheServiceServer) GetServerLoad(ctx context.Context, req *emptypb.Empty) (*pb.GetServerLoadResponse, error) {
	serverLoads, err := c.service.GetServerLoad(ctx)
	if err != nil {
		c.logger.Error("failed to fetch server load list", zap.Error(err))
		return &pb.GetServerLoadResponse{
			StatusResponse: &pb.CommonResponse{
				Message: "failed to fetch server load list",
				Code:    int32(codes.Internal),
				Success: false,
			},
		}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pb.GetServerLoadResponse{
		StatusResponse: &pb.CommonResponse{
			Message: "server registered to cache",
			Code:    int32(codes.OK),
			Success: true,
		},
		ServerLoad: serverLoads,
	}, nil
}

func (c *CacheServiceServer) GetUsersByChatId(ctx context.Context, req *pb.ChatDetails) (*pb.UsersByChatIdResponse, error) {
	users, err := c.service.GetUsersByChatId(ctx, req.ChatId)
	if err != nil {
		c.logger.Error("failed to fetch users list", zap.Error(err))
		return &pb.UsersByChatIdResponse{
			StatusResponse: &pb.CommonResponse{
				Message: "failed to fetch users list",
				Code:    int32(codes.Internal),
				Success: false,
			},
		}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pb.UsersByChatIdResponse{
		StatusResponse: &pb.CommonResponse{
			Message: "server registered to cache",
			Code:    int32(codes.OK),
			Success: true,
		},
		UserList: users,
	}, nil
}

func (c *CacheServiceServer) RegisterUserToCache(ctx context.Context, req *pb.UserDetails) (*pb.CommonResponse, error) {
	err := c.service.AddUserToCache(ctx, req)
	if err != nil {
		c.logger.Error("failed to register user to cache", zap.Error(err))
		return &pb.CommonResponse{
			Message: "failed to register user to cache",
			Code:    int32(codes.Internal),
			Success: false,
		}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pb.CommonResponse{
		Message: "user registered to cache",
		Code:    int32(codes.OK),
		Success: true,
	}, nil
}

func (c *CacheServiceServer) UnRegisterUserFromCache(ctx context.Context, req *pb.UserDetails) (*pb.CommonResponse, error) {
	err := c.service.RemoveUserFromCache(ctx, req.UserId)
	if err != nil {
		c.logger.Error("failed to yeet user from cache", zap.Error(err))
		return &pb.CommonResponse{
			Message: "failed to yeet user from cache",
			Code:    int32(codes.Internal),
			Success: false,
		}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pb.CommonResponse{
		Message: "user yeeted from cache",
		Code:    int32(codes.OK),
		Success: true,
	}, nil
}

func (c *CacheServiceServer) SetServerLoad(ctx context.Context, req *pb.SetServerLoadRequest) (*pb.CommonResponse, error) {
	err := c.service.SetServerLoad(ctx, req.ServerLoad)
	if err != nil {
		c.logger.Error("failed to update load for the servers", zap.Error(err))
		return &pb.CommonResponse{
			Message: "failed to update load for the servers",
			Code:    int32(codes.Internal),
			Success: false,
		}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pb.CommonResponse{
		Message: "server load counts updated",
		Code:    int32(codes.OK),
		Success: true,
	}, nil
}

func (c *CacheServiceServer) RegisterUserToChat(ctx context.Context, req *pb.UserToChat) (*pb.CommonResponse, error) {
	err := c.service.AddUserToChat(ctx, req.User.UserId, req.Chat.ChatId)
	if err != nil {
		c.logger.Error("failed to register user to cache", zap.Error(err))
		return &pb.CommonResponse{
			Message: "Failed to register user to cache",
			Code:    int32(codes.Internal),
			Success: false,
		}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pb.CommonResponse{
		Message: "user registered to cache",
		Code:    int32(codes.OK),
		Success: true,
	}, nil
}

func (c *CacheServiceServer) UnRegisterUserFromChat(ctx context.Context, req *pb.UserToChat) (*pb.CommonResponse, error) {
	err := c.service.RemoveUserFromChat(ctx, req.User.UserId, req.Chat.ChatId)
	if err != nil {
		c.logger.Error("failed to yeet user from cache", zap.Error(err))
		return &pb.CommonResponse{
			Message: "failed to yeet user from cache",
			Code:    int32(codes.Internal),
			Success: false,
		}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pb.CommonResponse{
		Message: "user yeeted from cache",
		Code:    int32(codes.OK),
		Success: true,
	}, nil
}
