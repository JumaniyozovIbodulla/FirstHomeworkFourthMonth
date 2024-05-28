package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	f "hm/files"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	f.UnimplementedFilesServer
}

func (s *server) CreateFile(ctx context.Context, req *f.CreateRequest) (*f.CreateResponse, error) {
	fmt.Println("Came to CreateFile: ", req)
	_, err := os.Create(req.FileName)
	if err != nil {
		log.Printf("failed to create file: %v", err)
		return &f.CreateResponse{}, err
	}
	path, err := os.Getwd()

	if err != nil {
		log.Printf("failed to get file path: %v", err)
		return &f.CreateResponse{}, err
	}
	return &f.CreateResponse{Path: path + "/" + req.FileName}, nil
}

func (s *server) IsFileExists(ctx context.Context, req *f.CreateRequest) (*f.CreateResponse, error) {
	fmt.Println("Came to IsFileExists: ", req)
	_, err := os.Stat(req.FileName)
	if err != nil {
		log.Printf("%s is not exists", req.FileName)
		return &f.CreateResponse{}, err
	}
	return &f.CreateResponse{IsExists: true}, nil
}

func (s *server) WroteToFile(ctx context.Context, req *f.CreateRequest) (*f.CreateResponse, error) {
	fmt.Println("Came to WroteToFile: ", req)
	file, err := os.Create(req.FileName)
	if err != nil {
		log.Printf("failed to create file: %v", err)
		return &f.CreateResponse{}, err
	}
	_, err = file.WriteString(req.Message)

	if err != nil {
		log.Printf("failed to write the file: %v", err)
		return &f.CreateResponse{}, err
	}
	return &f.CreateResponse{IsExists: true, IsWrote: true}, nil
}


func (s *server) ReadFromFile(ctx context.Context, req *f.CreateRequest) (*f.CreateResponse, error) {
	fmt.Println("Came to ReadFromFile: ", req)
	byteData, err := os.ReadFile(req.FileName)
	if err != nil {
		log.Printf("failed to read the file: %v", err)
		return &f.CreateResponse{}, err
	}
	convertedString := string(byteData)

	return &f.CreateResponse{Message: convertedString}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	f.RegisterFilesServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
