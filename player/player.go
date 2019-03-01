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
const (
	reportFrequency int = 10000
)

func Start(name string, address string){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewDirectorClient(conn)


	registerCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	p := pb.Player{Name:name}
	r, err := c.RegisterPlayer(registerCtx, &p)

	removeCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer c.RemovePlayer(removeCtx, &p) //remove from store
	fmt.Println(r)
	fmt.Println(err)


	playersCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	allPlayers, err := c.GetPlayers(playersCtx,&pb.Filter{})
	fmt.Println(allPlayers)
	fmt.Println(err)


}