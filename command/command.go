package command

import (
	"context"
	"fmt"
	directorProto "github.com/jwprillaman/orchestra/director/proto"
	playerProto "github.com/jwprillaman/orchestra/player/proto"
	"google.golang.org/grpc"
	"log"
	"strings"
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
		p, err := c.GetPlayers(ctx, &directorProto.Filter{PlayerName: "", SongIds: ""})
		defer cancel()

		if err != nil {
			panic(err)
		}
		fmt.Println(p)

	case "song":
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v\n", err)
		}
		defer conn.Close()
		c := playerProto.NewPlayerClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)

		split := strings.Split(input, " ")
		name := split[0]
		params := split[1:]

		r, err := c.Play(ctx, &playerProto.PlayRequest{Name:name, Params:params})
		defer cancel()

		if err != nil {
			panic(err)
		}
		fmt.Println(r)

	default:
		fmt.Printf("Invalid command : %v\n")
	}
}
