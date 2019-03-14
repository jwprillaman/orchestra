# Orchestra
## Author : James Prillaman

A tool to manage threads and system information across multiple machines. A director servie will manage and track multiple players.

### Definitions

#### Director
A single service that players will register with. A director will be responsible for tracking memory usage and deploying songs to players.

#### Player
A player is a client that runs on one or multiple machines in a cluster. It registers with a director and allows a director to deploy songs on it.

#### Song
A song is a program that is given to the director to be played on a registered player. Currently the player choosen to accept and run the song is choosen to be that with the most available memory.



### Quick Start

Run a director
```
orchestra -s director -p 50001
```

Run player
```
orchestra -s player -a localhost:50002
```

Run song
```
orchestra -s song -a localhost:50001 -i "/bin/bash orchestra/example/song.sh"
```

