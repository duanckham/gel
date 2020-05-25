package agent

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"github.com/duanckham/gel/gel"
	"github.com/duanckham/gel/pb"
)

var client pb.GelServiceClient

// New return an agent client.
func New(serverHost string, serverPort int32, period time.Duration) gel.Gel {
	c, err := grpc.Dial(fmt.Sprintf("%s:%d", serverHost, serverPort), grpc.WithInsecure())
	if err != nil {
		// TODO
		fmt.Println("* grpc.Dial err:", err)
	}

	// Connect to server.
	client = pb.NewGelServiceClient(c)

	// Start to collect data.
	return runGelAgent(context.Background(), period)
}

func runGelAgent(ctx context.Context, period time.Duration) gel.Gel {
	g := gel.New(period)

	g.SetTrigger(func(r *pb.Record) {
		_, err := client.SyncRecord(ctx, r)
		if err != nil {
			fmt.Println("* grpc.SyncRecord err:", err)
		}
	})

	return g
}
