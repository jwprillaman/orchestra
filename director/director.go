package director

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/jwprillaman/orchestra/director/proto"
	"github.com/jwprillaman/orchestra/song"
	"google.golang.org/grpc"
	"log"
	"net"
	"reflect"
	"sync"
	"time"
)

type playerStore struct {
	store map[string]Report
	mux   sync.Mutex
}

type server struct{}

type Report struct {
	Name     string
	TotalRam uint64
	FreeRam  uint64
	SongIds  []int64
}

var allPlayers = &playerStore{make(map[string]Report), sync.Mutex{}}

func (r *Report) Set(report *pb.PlayerReport) {
	r.Name = report.Name
	r.TotalRam = report.TotalRam
	r.FreeRam = report.FreeRam
	r.SongIds = report.SongIds
}

//get player with most memory available
func (ps *playerStore) GetBest() (string, error) {
	ps.mux.Lock()
	output := ""
	var err error = nil
	var mostMemPlayer Report = Report{}
	for address, player := range ps.store {
		if reflect.DeepEqual(mostMemPlayer, Report{}) || player.FreeRam > mostMemPlayer.FreeRam {
			mostMemPlayer = player
			output = address
		}
	}

	ps.mux.Unlock()
	if output == "" {
		err = errors.New("No registerd players to play song")
	}
	return output, err
}

//list all players that are registered
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

//get all playing songs on all players
func (s *server) GetSongs(context context.Context, filter *pb.Filter) (*pb.Songs, error) {

	allPlayers.mux.Lock()

	songs := make([]int64, 0)
	players := make([]string, 0)

	for _, player := range allPlayers.store {
		for _, songId := range player.SongIds {
			songs = append(songs, songId)
			players = append(players, player.Name)
		}
	}

	allPlayers.mux.Unlock()

	return &pb.Songs{Ids: songs, Players: players}, nil
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
		allPlayers.store[input.Name] = Report{}
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

//report player stats to be stored with director
func (s *server) Report(context context.Context, report *pb.PlayerReport) (*pb.Response, error) {
	fmt.Println(report)
	fmt.Println("\tsongs : ", report.SongIds)
	success := true
	msg := "success"
	if _, exists := allPlayers.store[report.Name]; exists {
		newReport := Report{}
		newReport.Set(report)
		allPlayers.store[report.Name] = newReport
	} else {
		success = false
		msg = fmt.Sprintf("%v is not registered with director\n", report.Name)
	}

	return &pb.Response{Success: success, Msg: msg}, nil
}

func (s *server) GetBestPlayer(context context.Context, report *pb.Filter) (*pb.Player, error) {
	address, err := allPlayers.GetBest()
	return &pb.Player{Name: address}, err
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

func Play(address string, input string) error {
	//connect to director and create client
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	client := pb.NewDirectorClient(conn)

	//register player with director
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	player, err := client.GetBestPlayer(ctx, &pb.Filter{})

	if err != nil {
		return err
	} else {
		fmt.Println("got address : %v\n", address)
		song.Play(player.Name, input)
		return nil
	}
}
