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

	fmt.Println("service :", *service)
	fmt.Println("port :", *port)
	fmt.Println("address :", *address)
	fmt.Println("command :", *cmd)
	fmt.Println("input :", *input)

	if *cmd != "" && *address != "" {
		command.Run(*address, *cmd, *input)
	}

	switch *service {
	case "director":
		director.Start(*port)
	case "player":
		player.Start(*address, *port)
	case "song":
		song.Start(*address, *input)
	default:
		fmt.Println("invalid service : ", *service)
	}

}
