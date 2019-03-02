package player

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"os/signal"
	"runtime"
	"time"

	pb "github.com/jwprillaman/orchestra/director/proto"
)

type Model struct {
	Address string
	SongIds []int64
	Mem uint64
}

type Report struct {
	Mem uint64
}
const (
	reportFrequency int = 1
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

	if err != nil {
		panic(err)
	}
	if !r.Success {
		panic(fmt.Sprintf("Could not register %v", name))
	}

	ch := make(chan Report)
	defer cleanup(c,p,ch) //remove from store

	osCh := make(chan os.Signal, 1)
	signal.Notify(osCh, os.Interrupt)

	go func(){
		<-osCh
		cleanup(c,p,ch)

	}()



	go monitor(ch)

	for {
		select{
		case x := <- ch:
			ctx, _ := context.WithTimeout(context.Background(), time.Second)
			r,err := c.Report(ctx, &pb.PlayerReport{Name:name, Mem:x.Mem,SongIds:make([]int64, 0)})
			if err != nil || !r.Success{
				log.Fatal("Could not communicate with director")
			}
		}
	}

}

func monitor(ch chan Report) {
	memStats := runtime.MemStats{}
	for {
		runtime.ReadMemStats(&memStats)
		select {
		case ch <- Report{memStats.Sys}:
		case <-ch:
			return
		}
		time.Sleep(time.Duration(reportFrequency) * time.Second)
	}

}

func cleanup(c pb.DirectorClient, p pb.Player, ch chan Report){
	fmt.Printf("I am running cleanup!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n")
	close(ch) //close channel to stop goroutine

	removeCtx, _ := context.WithTimeout(context.Background(), time.Second)
	c.RemovePlayer(removeCtx, &p)
}