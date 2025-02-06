package main

import (
	"log"
	"net"

	"github.com/dharmasatrya/hotel-booking/user-service/config"

	"github.com/dharmasatrya/hotel-booking/user-service/service"

	"github.com/dharmasatrya/hotel-booking/user-service/repository"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	pb "github.com/dharmasatrya/hotel-booking/user-service/proto"
)

func main() {
	godotenv.Load()

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}

	grpcServer := grpc.NewServer()

	config.InitDatabase()

	db := config.InitDatabase()
	if err != nil {
		log.Fatalf("Error connecting to db")
	}

	userRepository := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepository)
	pb.RegisterUserServiceServer(grpcServer, userService)

	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
