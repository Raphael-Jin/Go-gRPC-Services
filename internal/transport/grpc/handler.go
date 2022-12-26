package grpc

import (
	"context"
	"log"
	"net"

	"github.com/Raphael-Jin/Go-gRPC-Services/internal/rocket"
	rkt "github.com/Raphael-Jin/Go-gRPC-Services/protos/rocket/v1"
	"google.golang.org/grpc"
)

type RocketService interface {
	GetRocketByID(ctx context.Context, id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}

type Handler struct {
	RocketService RocketService
}

func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Print("Could not listen on port 50051")
		return err
	}
	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("fail to serve: %s\n", err)
		return err
	}
	return nil
}

func (h Handler) GetRocket(ctx context.Context, req *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	return &rkt.GetRocketResponse{}, nil
}

func (h Handler) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	return &rkt.AddRocketResponse{}, nil
}

func (h Handler) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	return &rkt.DeleteRocketResponse{}, nil
}
