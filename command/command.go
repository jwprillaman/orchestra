package command

import (
	"context"
	"fmt"
	pb "github.com/jwprillaman/orchestra/director/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Run(address string, input string) {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewDirectorClient(conn)

	switch input {
	case "players":
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		p, err := c.GetPlayers(ctx, &pb.Filter{PlayerName: "", SongIds: ""})
		defer cancel()

		if err != nil {
			panic(err)
		}
		fmt.Println(p)

	default:
		fmt.Printf("Invalid command : %v\n")
	}
}
