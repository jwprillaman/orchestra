package player

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "github.com/jwprillaman/orchestra/director/directorproto"
)

func Start(address string){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewDirectorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()


	r, err := c.Report(ctx, &pb.PlayerReport{Address:"player1", Mem:55, SongIds:[]int64{0,1,2}})
	fmt.Println(r)
	fmt.Println(err)
}