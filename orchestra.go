package main

import (
	"flag"
	"fmt"
	"github.com/jwprillaman/orchestra/command"
	"github.com/jwprillaman/orchestra/director"
	"github.com/jwprillaman/orchestra/player"
	"github.com/jwprillaman/orchestra/song"
)

func main() {
	service := flag.String("s", "", "Service")
	port := flag.Int("p", -1, "Port")
	address := flag.String("a", "", "Address")
	cmd := flag.String("c", "", "Command")
	input := flag.String("i", "", "Command Input")

	flag.Parse()

	if *cmd != "" && *address != "" {
		command.Run(*address, *cmd, *input)
	}

	switch *service {
	case "director":
		if *port == -1 {
			panic("Director service requires a valid port")
		}
		director.Start(*port)
	case "player":
		if *port == -1 || *address == "" {
			panic("Player service requires a valid port and director address")
		}
		player.Start(*address, *port)
	case "song":
		if *address == "" || *input == "" {
			panic("Song service requires a valid input and player address")
		}
		song.Start(*address, *input)
	default:
		fmt.Println("invalid service : ", *service)
	}

}
