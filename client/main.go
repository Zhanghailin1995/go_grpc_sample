package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "hello/proto"

	"google.golang.org/grpc"
)

const (
	defaultName = "world!"
)

var (
	addr = flag.String("addr", "127.0.0.1:19992", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterRequestServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.XCall(ctx, &pb.GreeterRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	for i := 0; i < 10; i++ {
		start := time.Now()
		XCall(ctx, &c, name)
		end := time.Now()
		log.Printf("cost %d", end.UnixMilli()-start.UnixMilli())
	}

}

func XCall(ctx context.Context, c *pb.GreeterRequestServiceClient, name *string) {
	r, err := (*c).XCall(ctx, &pb.GreeterRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
