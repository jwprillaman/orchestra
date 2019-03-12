package song

import (
	"errors"
	"fmt"
	"github.com/jwprillaman/orchestra/song/grpc"
	"strings"
)

const (
	communication_implementation = "grpc"
)

func Play(address string, input string) error {
	//parse input
	split := strings.Split(input, " ")
	name := split[0]
	params := split[1:]

	switch communication_implementation {
	case "grpc":
		return grpc.Play(address, name, params)
	default:
		return errors.New(fmt.Sprintf("%v communication not supported", communication_implementation))
	}
}
