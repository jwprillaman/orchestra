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
	name := flag.String("n", "", "Name")
	cmd := flag.String("c", "", "Command")

	flag.Parse()

	fmt.Println("name : " + *name)
	fmt.Println("service :", *service)
	fmt.Println("port :", *port)
	fmt.Println("address :", *address)

	if *cmd != "" && *address != "" {
		command.Run(*address, *cmd)
	}

	switch *service {
	case "director":
		director.Start(*port)
	case "player":
		player.Start(*name, *address)
	default:
		fmt.Println("invalid service : ", *service)
	}

}
