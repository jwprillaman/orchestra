package player

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "github.com/jwprillaman/orchestra/director/proto"
)

type Model struct {
	Address string
	SongIds []string
}

func Start(name string, address string){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewDirectorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()


	p := pb.Player{Name:name}
	r, err := c.RegisterPlayer(ctx, &p)
	defer c.RemovePlayer(ctx, &p) //remove from store
	fmt.Println(r)
	fmt.Println(err)

	allPlayers, err := c.GetPlayers(ctx,&pb.Filter{})
	fmt.Println(allPlayers)
	fmt.Println(err)


}