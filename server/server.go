package server

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/duanckham/gel/gel"
	"github.com/duanckham/gel/pb"
	"github.com/golang/protobuf/ptypes/empty"
)

// GelServer ...
type GelServer struct{}

// New return a server.
func New(port int32) {
	s, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		// TODO
		fmt.Println("* net.Listen err:", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGelServiceServer(grpcServer, &GelServer{})

	grpcServer.Serve(s)
}

// SyncRecord endpoint receive agent.
func (gs *GelServer) SyncRecord(ctx context.Context, in *pb.Record) (*empty.Empty, error) {
	reader, done := gel.Read(in)
	start := time.Now()

	go func() {
		for {
			select {
			case data := <-reader:
				switch data.T {
				case "log":
					fmt.Println("* (log)", data.D, data.V)

				case "number", "instant":
					fmt.Println("* (number or instant)", data.T, data.D, data.V)
				}
			case <-done:
				fmt.Println("* data all processed, cost:", time.Since(start))
				return
			}
		}
	}()

	return &empty.Empty{}, nil
}
