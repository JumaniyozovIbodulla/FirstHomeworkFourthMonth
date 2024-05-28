package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	f "hm/files"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := f.NewFilesClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)

	defer cancel()

	r, err := c.ReadFromFile(ctx, &f.CreateRequest{FileName: "check.txt"})
	if err != nil {
		log.Fatalf("could not create a file: %v", err)
	}
	fmt.Println(r.Message)
}
