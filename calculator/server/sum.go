package main

import (
	"context"
	"log"

	pb "grpc-go-project/calculator/proto"
)

func (s *Server) Sum(stx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
