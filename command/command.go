package command

import (
	"context"
	"fmt"
	directorProto "github.com/jwprillaman/orchestra/director/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Run(address string, command string, input string) {

	switch command {
	case "players":
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v\n", err)
		}
		defer conn.Close()
		c := directorProto.NewDirectorClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		p, err := c.GetPlayers(ctx, &directorProto.Filter{})
		defer cancel()

		if err != nil {
			panic(err)
		}
		fmt.Println(p)
	case "songs":
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v\n", err)
		}
		defer conn.Close()
		c := directorProto.NewDirectorClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		s, err := c.GetSongs(ctx, &directorProto.Filter{})
		defer cancel()

		if err != nil {
			panic(err)
		}
		fmt.Println(s)
	default:
		fmt.Printf("Invalid command : %v\n")
	}
}
