package server

import (
	"context"

	api "github.com/ozonva/ova-reason-api/pkg/ova-reason-api"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ReasonServer struct {
	api.UnimplementedReasonRpcServer
	logger *zerolog.Logger
}

func NewReasonRpcServer(logger *zerolog.Logger) api.ReasonRpcServer {
	return &ReasonServer{
		UnimplementedReasonRpcServer: api.UnimplementedReasonRpcServer{},
		logger:                       logger,
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
	return s.UnimplementedReasonRpcServer.ListReasons(context, empty)
}

func (s *ReasonServer) RemoveReason(context context.Context, request *api.RemoveReasonRequest) (*emptypb.Empty, error) {
	s.logger.Info().Msgf("RemoveReason request: %v", request)
	return s.UnimplementedReasonRpcServer.RemoveReason(context, request)
}
