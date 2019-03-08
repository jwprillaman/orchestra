package player

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"sync"
	"time"

	directorProto "github.com/jwprillaman/orchestra/director/proto"
	playerProto "github.com/jwprillaman/orchestra/player/proto"
)

const (
	reportFrequency int = 1
)

var playerName string
var allSongs = songStore{make(map[int64]struct{}), sync.Mutex{}, make([]int64, 0)}

type songStore struct {
	store map[int64]struct{}
	mux   sync.Mutex
	ids   []int64
}

func (s *songStore) Add(key int64) {
	s.mux.Lock()
	if _, exists := s.store[key]; !exists {
		s.store[key] = struct{}{}
		s.ids = append(s.ids, key)
	}
	s.mux.Unlock()
}

func (s *songStore) Remove(key int64) {
	s.mux.Lock()
	if _, exists := s.store[key]; exists {
		delete(s.store, key)
		for i := range s.ids {
			if s.ids[i] == key {
				s.ids[i] = s.ids[len(s.ids)-1]
			}
		}
		s.ids = s.ids[:len(s.ids)-1]
	}
	s.mux.Unlock()
}

func (s *songStore) Contains(key int64) bool {
	s.mux.Lock()
	_, output := s.store[key]
	s.mux.Unlock()
	return output
}

type server struct{}

func (*server) Play(ctx context.Context, req *playerProto.PlayRequest) (*playerProto.PlayResponse, error) {
	cmd := exec.Command(req.Name, req.Params...)
	//get std and std out
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	//add to songids
	allSongs.Add(int64(cmd.Process.Pid))
	//wait until finished and remove from songids
	go func() {
		eo, _ := ioutil.ReadAll(stdout)
		so, _ := ioutil.ReadAll(stderr)
		fmt.Println(string(eo))
		fmt.Println(string(so))

		processState, err := cmd.Process.Wait()
		fmt.Println(err)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(processState.Exited())
		fmt.Println(processState.Success())
		fmt.Println(processState.String())
		allSongs.Remove(int64(cmd.Process.Pid))
	}()
	return &playerProto.PlayResponse{}, nil
}

func (*server) Stop(ctx context.Context, song *playerProto.StopRequest) (*playerProto.StopResponse, error) {
	return &playerProto.StopResponse{}, nil
}

//Register player and begin send stats to director
func Start(directorAddress string, port int) {
	playerName = fmt.Sprintf("localhost:%v", port)

	//create player service
	fmt.Printf("starting player at : %v\n", port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("failed to listen on : %v\n", err)
	}
	s := grpc.NewServer()
	playerProto.RegisterPlayerServer(s, &server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	//connect to director and create client
	conn, err := grpc.Dial(directorAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	client := directorProto.NewDirectorClient(conn)

	//register player with director
	registerCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	player := directorProto.Player{Name: playerName}
	r, err := client.RegisterPlayer(registerCtx, &player)
	if err != nil {
		panic(err)
	}
	if !r.Success {
		panic(fmt.Sprintf("Could not register %v", playerName))
	}
	//Create reporting channel
	ch := make(chan directorProto.PlayerReport)
	//cleanup connection, client, and channel
	defer cleanup(conn, client, player, ch) //remove from store

	//listen for interruption and cleanup if found
	osCh := make(chan os.Signal, 1)
	signal.Notify(osCh, os.Interrupt)
	go func() {
		<-osCh
		cleanup(conn, client, player, ch)

	}()

	//get stats at interval
	go monitor(ch)

	//collect stats and send to director
	for {
		select {
		case x := <-ch:
			ctx, _ := context.WithTimeout(context.Background(), time.Second)
			r, err := client.Report(ctx, &directorProto.PlayerReport{Name: playerName, Alloc: x.Alloc, TotalAlloc: x.TotalAlloc, Sys: x.Sys, Mallocs: x.Mallocs, Frees: x.Frees, LiveObjects: x.LiveObjects, PauseTotalNs: x.PauseTotalNs, NumGC: x.NumGC, SongIds: allSongs.ids})
			if err != nil || !r.Success {
				log.Fatal("Could not communicate with director")
			}
		}
	}
}

//Monitor stats and send to channel at interval
func monitor(ch chan directorProto.PlayerReport) {
	stats := runtime.MemStats{}
	for {
		runtime.ReadMemStats(&stats)
		select {
		case ch <- directorProto.PlayerReport{Name: playerName, Alloc: stats.Alloc, TotalAlloc: stats.TotalAlloc, Sys: stats.Sys, Mallocs: stats.Mallocs, Frees: stats.Frees, LiveObjects: stats.Mallocs - stats.Frees, PauseTotalNs: stats.PauseTotalNs, NumGC: stats.NumGC, SongIds: allSongs.ids}:
		case <-ch:
			return
		}
		time.Sleep(time.Duration(reportFrequency) * time.Second)
	}

}

//cleanup connection, client, and report channel
func cleanup(connection *grpc.ClientConn, client directorProto.DirectorClient, p directorProto.Player, playerReportChannel chan directorProto.PlayerReport) {
	//close channel for goroutine collecting report info
	close(playerReportChannel)

	//Unregister player from director
	removeCtx, _ := context.WithTimeout(context.Background(), time.Second)
	r, err := client.RemovePlayer(removeCtx, &p)
	if err != nil || !r.Success {
		//TODO handle this with retries
	}
	//Close connection to director
	connection.Close()
}
