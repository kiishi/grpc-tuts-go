package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/deolu-asenuga/learn-grpc/point"
)

type RouteServer struct{}

func (this *RouteServer) GetLocation(ctx context.Context, req *point.Point) (*point.PointResponse, error) {
	log.Println(req.X, req)

	return &point.PointResponse{Message: "Worked ", Code: 500}, nil
}

func Start() {
	srv := grpc.NewServer()
	point.RegisterRouteGuideServer(srv, &RouteServer{})
	listerner, err := net.Listen("tcp", ":8888")
	err = srv.Serve(listerner)
	if err != nil {
		log.Fatalf("An error occured : %v", err)
	}
}
