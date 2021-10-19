package level

import (
	"context"
	"log"
	"time"

	"github.com/Gavin-Lau/hi-level/protocol/LEVELCONFIG"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func GetConfigByKeyClient(keys []string) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := LEVELCONFIG.NewGConfigClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for _, key := range keys {
		r, err := c.GetConfigByKey(ctx, &LEVELCONFIG.GConfigReq{Key: []byte(key)})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Value)
	}
	return nil
}
