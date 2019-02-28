package main

import (
	"flag"
	"fmt"
	"github.com/jwprillaman/orchestra/director"
	"github.com/jwprillaman/orchestra/player"
)

func main() {
	service := flag.String("service", "director", "service to start")
	port := flag.Int("port", 50001, "port for service")
	address := flag.String("address", "localhost:50001", "address of director")

	flag.Parse()

	fmt.Println("service :", *service)
	fmt.Println("port :", *port)
	fmt.Println("address :", *address)

	switch *service {
	case "director":
		director.Start(*port)
	case "player":
		player.Start(*address)
	default:
		fmt.Println("invalid service : ", *service)
	}

}
