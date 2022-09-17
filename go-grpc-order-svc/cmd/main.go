package main

import (
	"fmt"
	"github.com/Tushar456/go-grpc-order-svc/pkg/client"
	"github.com/Tushar456/go-grpc-order-svc/pkg/config"
	"github.com/Tushar456/go-grpc-order-svc/pkg/db"
	"github.com/Tushar456/go-grpc-order-svc/pkg/pb"
	"github.com/Tushar456/go-grpc-order-svc/pkg/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product Svc on", c.Port)
	productSvc := client.InitProductServiceClient(c.ProductSvcUrl)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}
	fmt.Println("Order Svc on", c.Port)
	s := services.Server{
		H:          h,
		ProductSvc: productSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
