// Code generated by sqlc-grpc (https://github.com/walterwanderley/sqlc-grpc). DO NOT EDIT.

package pguuid

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "uuidcheck/api/pguuid/v1"
	"uuidcheck/internal/validation"
)

type Service struct {
	pb.UnimplementedPguuidServiceServer
	querier *Queries
}

func (s *Service) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var arg CreateProductParams
	arg.ID = req.GetId()
	if v := req.GetCategory(); v != nil {
		arg.Category = pgtype.Int4{Valid: true, Int32: v.Value}
	}

	result, err := s.querier.CreateProduct(ctx, arg)
	if err != nil {
		slog.Error("CreateProduct sql call failed", "error", err)
		return nil, err
	}
	return &pb.CreateProductResponse{Value: result}, nil
}

func (s *Service) CreateProductReturnAll(ctx context.Context, req *pb.CreateProductReturnAllRequest) (*pb.CreateProductReturnAllResponse, error) {
	var arg CreateProductReturnAllParams
	arg.ID = req.GetId()
	if v := req.GetCategory(); v != nil {
		arg.Category = pgtype.Int4{Valid: true, Int32: v.Value}
	}

	result, err := s.querier.CreateProductReturnAll(ctx, arg)
	if err != nil {
		slog.Error("CreateProductReturnAll sql call failed", "error", err)
		return nil, err
	}
	return &pb.CreateProductReturnAllResponse{Product: toProduct(result)}, nil
}

func (s *Service) CreateProductReturnPartial(ctx context.Context, req *pb.CreateProductReturnPartialRequest) (*pb.CreateProductReturnPartialResponse, error) {
	var arg CreateProductReturnPartialParams
	arg.ID = req.GetId()
	if v := req.GetCategory(); v != nil {
		arg.Category = pgtype.Int4{Valid: true, Int32: v.Value}
	}

	result, err := s.querier.CreateProductReturnPartial(ctx, arg)
	if err != nil {
		slog.Error("CreateProductReturnPartial sql call failed", "error", err)
		return nil, err
	}
	return &pb.CreateProductReturnPartialResponse{CreateProductReturnPartialRow: toCreateProductReturnPartialRow(result)}, nil
}

func (s *Service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var arg CreateUserParams
	if v := req.GetId(); v != nil {
		if err := json.Unmarshal([]byte(v.GetValue()), &arg.ID); err != nil {
			err = fmt.Errorf("invalid ID: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
	}
	if v := req.GetLocation(); v != nil {
		if err := json.Unmarshal([]byte(v.GetValue()), &arg.Location); err != nil {
			err = fmt.Errorf("invalid Location: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
	}

	result, err := s.querier.CreateUser(ctx, arg)
	if err != nil {
		slog.Error("CreateUser sql call failed", "error", err)
		return nil, err
	}
	if uuidStr, err := result.MarshalJSON(); err != nil {
		return nil, fmt.Errorf("failed to convert UUID to string: %w", err)
	} else {
		return &pb.CreateUserResponse{Value: wrapperspb.String(string(uuidStr))}, nil
	}
}

func (s *Service) CreateUserReturnAll(ctx context.Context, req *pb.CreateUserReturnAllRequest) (*pb.CreateUserReturnAllResponse, error) {
	var arg CreateUserReturnAllParams
	if v := req.GetId(); v != nil {
		if err := json.Unmarshal([]byte(v.GetValue()), &arg.ID); err != nil {
			err = fmt.Errorf("invalid ID: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
	}
	if v := req.GetLocation(); v != nil {
		if err := json.Unmarshal([]byte(v.GetValue()), &arg.Location); err != nil {
			err = fmt.Errorf("invalid Location: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
	}

	result, err := s.querier.CreateUserReturnAll(ctx, arg)
	if err != nil {
		slog.Error("CreateUserReturnAll sql call failed", "error", err)
		return nil, err
	}
	return &pb.CreateUserReturnAllResponse{User: toUser(result)}, nil
}

func (s *Service) CreateUserReturnPartial(ctx context.Context, req *pb.CreateUserReturnPartialRequest) (*pb.CreateUserReturnPartialResponse, error) {
	var arg CreateUserReturnPartialParams
	if v := req.GetId(); v != nil {
		if err := json.Unmarshal([]byte(v.GetValue()), &arg.ID); err != nil {
			err = fmt.Errorf("invalid ID: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
	}
	if v := req.GetLocation(); v != nil {
		if err := json.Unmarshal([]byte(v.GetValue()), &arg.Location); err != nil {
			err = fmt.Errorf("invalid Location: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
	}

	result, err := s.querier.CreateUserReturnPartial(ctx, arg)
	if err != nil {
		slog.Error("CreateUserReturnPartial sql call failed", "error", err)
		return nil, err
	}
	return &pb.CreateUserReturnPartialResponse{CreateUserReturnPartialRow: toCreateUserReturnPartialRow(result)}, nil
}

func (s *Service) WithTx(tx pgx.Tx) *Service {
	return &Service{
		querier: s.querier.WithTx(tx),
	}
}
