package main

import (
	"context"
	pb "example.com/pb_source_files"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	defaultMessage       = "Nice to meet you too.."
	defaultRemmittedDate = "01-04-2023"
)

var (
	port          = flag.Int("port", 50051, "The server port")
	message       = flag.String("message", defaultMessage, "Hello nice to meet you...")
	remmittedDate = flag.String("empRemmittedDate", defaultRemmittedDate, "Salary remmitted date")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Salary status for : %v", in.GetName())
	//var pName = in.GetName()
	//_empName := &pName
	// var pEmpSal = in.GetEmpSal()
	// _empSal := &pEmpSal
	// var pRemmittdDate = in.GetRemmittedDate()
	// _remmittedDate := &pRemmittdDate

	return &pb.HelloReply{Message: message}, nil
}

func main() {
	fmt.Println("This is server")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	time.Sleep(2 * time.Second)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
