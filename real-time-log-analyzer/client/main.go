package main

import (
	"context"
	"log"
	pb "real-time-log-analyzer/api"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1. Connect to the Server
	conn, err := grpc.NewClient("localhost:50021", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error closing connection: %v", err)
		}
	}(conn) // Close when done

	// 2. Create a client from code
	c := pb.NewLogServiceClient(conn)

	// 3. Send Request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 4. Call the remote function "IngestLog" like local function
	r, err := c.IngestLog(ctx, &pb.LogEntry{
		ServiceName: "payment-service",
		Level:       "ERROR",
		Message:     "Database connection Timeout",
		Timestamp:   time.Now().Unix(),
	})
	if err != nil {
		log.Fatalf("Cannot ingest log: %v", err)
	}

	log.Printf("Server response: Success=%v\n", r.Success)

}
