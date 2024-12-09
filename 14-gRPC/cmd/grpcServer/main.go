package main

import (
	"14-gRPC/internal/database"
	"14-gRPC/internal/pb"
	"14-gRPC/internal/service"
	"database/sql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	gprcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(gprcServer, categoryService)
	reflection.Register(gprcServer)

	lis, err := net.Listen("grpc", ":50051")
	if err != nil {
		// log.Fatalf("failed to listen: %v", err)
		panic(err)
	}
	if err := gprcServer.Serve(lis); err != nil {
		// log.Fatalf("failed to serve: %v", err)
		panic(err)
	}

}
