package main

import (
	"flag"
	"fmt"
	"github.com/jwprillaman/orchestra/command"
	"github.com/jwprillaman/orchestra/director"
	"github.com/jwprillaman/orchestra/player"
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
		if *input == "" || *address == "" {
			panic("Song service requires a valid input and director address")
		}
		err := director.Play(*address, *input)
		if err != nil {
			panic(err)
		}
	default:
		fmt.Println("invalid service : ", *service)
	}

}
