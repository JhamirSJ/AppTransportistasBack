package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"AppTransportistasBack/config"
	"AppTransportistasBack/despachopb"
	"AppTransportistasBack/handlers"
)

func main() {
	// Inicializar conexión a MySQL
	config.InitDB()

	// Escuchar en el puerto 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Error escuchando en el puerto 50051: %v", err)
	}

	// Crear servidor gRPC
	grpcServer := grpc.NewServer()

	// Registrar el servicio DespachoService
	despachoService := &handlers.DespachoServer{}
	despachopb.RegisterDespachoServiceServer(grpcServer, despachoService)

	log.Println("🚀 Servidor gRPC corriendo en el puerto 50051")

	// Iniciar el servidor
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Error al iniciar servidor gRPC: %v", err)
	}
}
