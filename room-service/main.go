package main

import (
	"log"
	"net"

	"github.com/dharmasatrya/hotel-booking/room-service/config"
	"github.com/dharmasatrya/hotel-booking/room-service/repository"
	"github.com/dharmasatrya/hotel-booking/room-service/service"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	pb "github.com/dharmasatrya/hotel-booking/room-service/proto"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	db := config.InitDatabase()
	if err != nil {
		log.Fatalf("Error connecting to db")
	}

	roomRepository := repository.NewRoomRepository(db)
	roomService := service.NewRoomService(roomRepository)

	pb.RegisterUserServiceServer(grpcServer, roomService)

	log.Println("Room service is running on port 50052...")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
