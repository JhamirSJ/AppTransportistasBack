package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"AppTransportistasBack/saludarpb"

	"google.golang.org/grpc"
)

type server struct {
	saludarpb.UnimplementedSaludarServiceServer
}

func (s *server) SayHello(ctx context.Context, req *saludarpb.HelloRequest) (*saludarpb.HelloResponse, error) {
	name := req.GetName()
	result := fmt.Sprintf("¡Hola, %s! \nEsto es una prueba de gRPC", name)
	return &saludarpb.HelloResponse{Message: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	saludarpb.RegisterSaludarServiceServer(grpcServer, &server{})

	log.Println("Servidor gRPC corriendo en el puerto 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
