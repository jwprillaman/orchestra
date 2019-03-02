package director

import (
	"context"
	"fmt"
	pb "github.com/jwprillaman/orchestra/director/proto"
	"github.com/jwprillaman/orchestra/player"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

type server struct{}

type playerStore struct {
	store map[string]*player.Model
	mux sync.Mutex
}

	var allPlayers = &playerStore{make(map[string]*player.Model),sync.Mutex{}}


func (s *server) GetPlayers(context context.Context, filter *pb.Filter) (*pb.Players, error) {
	log.Printf("Filter : %v\n", filter.PlayerName)
	allPlayers.mux.Lock()
	names := make([]string, len(allPlayers.store))
	i := 0
	for name := range allPlayers.store {
		names[i] = name
		i++
	}
	allPlayers.mux.Unlock()
	return &pb.Players{Names: names}, nil
}

func (s *server) GetSongs(context context.Context, filter *pb.Filter) (*pb.Songs, error) {
	log.Printf("Filter : %v\n", filter.SongIds)
	return &pb.Songs{Ids: []int64{0, 1}, Players: []string{"player1", "player2"}}, nil
}

func (s *server) RegisterPlayer(context context.Context, input *pb.Player) (*pb.Response, error) {
	allPlayers.mux.Lock()
	success := true
	msg := "success"
	_, exists := allPlayers.store[input.Name]
	if exists {
		success = false
		msg = "name already registered"
	} else {
		allPlayers.store[input.Name] = &player.Model{Address: input.Name, SongIds: make([]int64, 0), Mem:0}
	}
	allPlayers.mux.Unlock()
	return &pb.Response{Success: success, Msg: msg}, nil
}

func (s *server) RemovePlayer(context context.Context, player *pb.Player) (*pb.Response, error) {
	allPlayers.mux.Lock()
	delete(allPlayers.store, player.Name)

	allPlayers.mux.Unlock()
	return &pb.Response{Success: true, Msg: "success"}, nil
}

func (s *server) Report(context context.Context, report *pb.PlayerReport) (*pb.Response, error) {
	fmt.Printf("Player : %v\n\tMem : %v\n\tSongs : %v\n", report.Mem, report.Mem, len(report.SongIds))
	success := true
	msg := "success"
	if v, exists := allPlayers.store[report.Name]; exists {
		v.Mem = report.Mem
		v.SongIds = report.SongIds
	} else {
		success = false
		msg = fmt.Sprintf("%v is not registered with director\n", report.Name)
	}

	return &pb.Response{Success: success, Msg: msg}, nil
}

func Start(port int) {
	fmt.Printf("starting director at : %v\n", port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("failed to listen on : %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterDirectorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
