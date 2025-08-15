package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/mtsdalmolin/grpc-protobuf/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

func createCategory(client pb.CategoryServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	response, err := client.CreateCategory(ctx, &pb.CreateCategoryRequest{
		Name:        "Test Category",
		Description: "This is a test category",
	})

	if err != nil {
		log.Fatalf("could not create category: %v", err)
	}

	log.Printf("Category created: %v", response)
}

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewCategoryServiceClient(conn)

	// createCategory(client)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Call ListCategories with proper parameters
	response, err := client.ListCategories(ctx, &pb.Blank{})
	if err != nil {
		log.Fatalf("could not list categories: %v", err)
	}

	log.Printf("Categories received: %v", response)

	// Print each category
	for i, category := range response.Categories {
		log.Printf("Category %d: ID=%s, Name=%s, Description=%s",
			i+1, category.Id, category.Name, category.Description)
	}
}
