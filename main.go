package main

// "context" // Removed unused import
import (
	"feed-service/graph"
	"feed-service/graph/generated"
	"feed-service/postservice"

	// pb "feed-service/proto/post"
	pb "feed-service/proto"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	// Start gRPC server
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		pb.RegisterPostServiceServer(grpcServer, &postservice.PostServiceServer{})
		reflection.Register(grpcServer)
		log.Println("gRPC server running on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Connect to gRPC client
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to gRPC: %v", err)
	}
	defer conn.Close()

	client := pb.NewPostServiceClient(conn)

	// Start GraphQL server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{Resolvers: &graph.Resolver{PostsClient: client}},
	))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("GraphQL server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
