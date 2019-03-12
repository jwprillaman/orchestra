package grpc

import (
	"context"
	"fmt"
	playerProto "github.com/jwprillaman/orchestra/player/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Play(address string, name string, params []string) error {

	//Dial player at address
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()

	//Create client
	c := playerProto.NewPlayerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	//Make request to player for input
	r, err := c.Play(ctx, &playerProto.PlayRequest{Name: name, Params: params})
	defer cancel()

	if err != nil {
		panic(err)
	}

	fmt.Println(r)
	return nil
}
