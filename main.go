package main

import (
	"aws_main/protoc/math_service"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// 서버를 위한 구조체
type server struct {
	math_service.UnimplementedMathServiceServer
}

// Add 메소드를 구현
func (s *server) Add(ctx context.Context, req *math_service.AddRequest) (*math_service.AddResponse, error) {
	result := req.GetA() + req.GetB()
	return &math_service.AddResponse{Sum: result}, nil
}

func main() {
	// 서버가 사용할 포트 설정
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	math_service.RegisterMathServiceServer(grpcServer, &server{})

	// gRPC 서버에서 Reflection을 사용할 수 있도록 설정 (클라이언트 개발 시 유용)
	reflection.Register(grpcServer)

	log.Printf("Server is listening on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
