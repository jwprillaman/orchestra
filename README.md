# Orchestra
## Author : James Prillaman

A tool to manage threads and system information across multiple machines. A director servie will manage and track multiple players.




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

