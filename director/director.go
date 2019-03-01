package director

import (
	"context"
	"fmt"
	pb "github.com/jwprillaman/orchestra/director/proto"
	"github.com/jwprillaman/orchestra/player"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

var playerStore map[string]*player.Model = make(map[string]*player.Model)

func (s *server) GetPlayers(context context.Context, filter *pb.Filter) (*pb.Players, error) {
	log.Printf("Filter : %v\n", filter.PlayerName)
	addresses := make([]string, len(playerStore))
	i := 0
	for address := range playerStore {
		addresses[i] = address
		i++
	}
	return &pb.Players{Addresses: addresses}, nil
}

func (s *server) GetSongs(context context.Context, filter *pb.Filter) (*pb.Songs, error) {
	log.Printf("Filter : %v\n", filter.SongIds)
	return &pb.Songs{Ids: []int64{0, 1}, Players: []string{"player1", "player2"}}, nil
}

func (s *server) Register(context context.Context, input *pb.Player) (*pb.Response, error) {
	success := true
	msg := "success"
	_, exists := playerStore[input.Address]
	if exists {
		success = false
		msg = "address already registered"
	} else {
		playerStore[input.Address] = &player.Model{Address: input.Address, SongIds: make([]string, 0)}
	}
	return &pb.Response{Success: success, Msg: msg}, nil
}

func (s *server) RemovePlayer(context context.Context, player *pb.Player) (*pb.Response, error) {
	return &pb.Response{Success: true, Msg: "success"}, nil
}

func (s *server) Report(context context.Context, report *pb.PlayerReport) (*pb.Response, error) {
	fmt.Printf("Player : %v\n\tMem : %v\n\tSongs : %v\n", report.Mem, report.Mem, len(report.SongIds))
	return &pb.Response{Success: true, Msg: "success"}, nil
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
