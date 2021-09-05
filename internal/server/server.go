package server

import (
	"context"
	"github.com/ozonva/ova-reason-api/internal/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"time"

	api "github.com/ozonva/ova-reason-api/pkg/ova-reason-api"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ReasonServer struct {
	api.UnimplementedReasonRpcServer
	logger     *zerolog.Logger
	reasonRepo repo.Repo
}

func NewReasonRpcServer(repo *repo.Repo) api.ReasonRpcServer {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	zLogger := zerolog.New(output).With().Timestamp().Logger()
	return &ReasonServer{
		UnimplementedReasonRpcServer: api.UnimplementedReasonRpcServer{},
		logger:                       &zLogger,
		reasonRepo:                   *repo,
	}
}

func (s *ReasonServer) CreateReason(context context.Context, request *api.CreateReasonRequest) (*api.CreateReasonResponse, error) {
	s.logger.Info().Msgf("CreateReason request: %v", request)
	return s.UnimplementedReasonRpcServer.CreateReason(context, request)
}

func (s *ReasonServer) DescribeReason(context context.Context, request *api.DescribeReasonRequest) (*api.DescribeReasonResponse, error) {
	s.logger.Info().Msgf("DescribeReason request: %v", request)
	return s.UnimplementedReasonRpcServer.DescribeReason(context, request)
}

func (s *ReasonServer) ListReasons(context context.Context, empty *emptypb.Empty) (*api.ListReasonsResponse, error) {
	s.logger.Info().Msgf("ListReasons request: %v", empty)

	result, err := s.reasonRepo.ListEntities(100, 0)
	if err != nil {
		s.logger.Error().Err(err).Msg("")
		return nil, status.Error(codes.Internal, "Internal")
	}

	list := make([]*api.Reason, 0)
	for _, v := range result {
		list = append(list, &api.Reason{
			Id:       v.Id,
			ActionId: v.ActionId,
			UserId:   v.UserId,
			Why:      v.Why,
		})
	}

	return &api.ListReasonsResponse{
		Items: list,
	}, nil
}

func (s *ReasonServer) RemoveReason(context context.Context, request *api.RemoveReasonRequest) (*emptypb.Empty, error) {
	s.logger.Info().Msgf("RemoveReason request: %v", request)
	return s.UnimplementedReasonRpcServer.RemoveReason(context, request)
}
