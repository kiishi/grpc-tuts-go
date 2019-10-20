package main

import (
	"context"
	"log"

	"github.com/deolu-asenuga/learn-grpc/point"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Printf(" An error occured : %v", err)
	}

	client := point.NewRouteGuideClient(conn)

	resp, err := client.GetLocation(context.Background(), &point.Point{X: 12, Y: 13})
	if err != nil {
		log.Printf(" An error occured : %v", err)
	}
	log.Println(resp)
}
