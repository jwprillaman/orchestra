package main

import (
	"flag"
	"fmt"
	"github.com/jwprillaman/orchestra/director"
	"github.com/jwprillaman/orchestra/player"
)

func main() {
	service := flag.String("s", "", "Service")
	port := flag.Int("p", -1, "Port")
	address := flag.String("a", "", "Address")
	name := flag.String("n", "", "Name")
	command := flag.String("c", "", "Command")

	flag.Parse()

	fmt.Println("name : " + *name)
	fmt.Println("service :", *service)
	fmt.Println("port :", *port)
	fmt.Println("address :", *address)

	switch *command {
	case "players":
		//TODO create command capability 
	default:
		fmt.Printf("%v is not a valid command\n", command)

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
