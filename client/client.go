package main

import (
	"context"
	"log"
	"time"

	"AppTransportistasBack/saludarpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	client := saludarpb.NewSaludarServiceClient(conn)

	req := &saludarpb.HelloRequest{
		Name: "Sebastian",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("Error llamando al servicio: %v", err)
	}

	log.Printf("Respuesta del servidor: %s", res.GetMessage())
}
