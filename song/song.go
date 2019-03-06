package song

import (
	"context"
	"fmt"
	playerProto "github.com/jwprillaman/orchestra/player/proto"
	"google.golang.org/grpc"
	"log"
	"strings"
	"time"
)

func Start(address string, input string) {

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

	r, err := c.Play(ctx, &playerProto.PlayRequest{Name: name, Params: params})
	defer cancel()

	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
