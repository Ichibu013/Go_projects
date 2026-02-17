package main

import (
	"context"
	"log"
	"net"
	pb "real-time-log-analyzer/api"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLogServiceServer // Required by gRPC for forward compatibility
}

// IngestLog implements log.LogServiceServer
// Notice the signature: It takes a Context and a LogEntry pointer.
func (s *server) IngestLog(_ context.Context, in *pb.LogEntry) (*pb.Ack, error) {
	// Simulate Processing
	log.Printf("Received: [%s] %s: %s", in.Level, in.ServiceName, in.Message)

	// Return Success
	return &pb.Ack{Success: true}, nil
}

func main() {
	// 1. Listen to TCP port
	lis, err := net.Listen("tcp", ":50021")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 2. Create a gRPC server object
	s := grpc.NewServer()

	// 3. Register server implementation
	pb.RegisterLogServiceServer(s, &server{})

	log.Printf("Starting gRPC server at %v\n", lis.Addr())

	// 4. Start Serving
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
