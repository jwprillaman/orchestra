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

func Start(address string){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewDirectorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()


	r, err := c.Register(ctx, &pb.Player{Address:"player1"})
	fmt.Println(r)
	fmt.Println(err)

	p, err := c.GetPlayers(ctx,&pb.Filter{})
	fmt.Println(p)
	fmt.Println(err)


}