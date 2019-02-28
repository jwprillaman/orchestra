# Orchestra
## Author : James Prillaman

A tool to manage threads and system information across multiple machines. A director servie will manage and track multiple players.




###Quick Start

Run a director
```
orchestra -service=director -port 50001
```

Run player
```
orchestra -service=player -address=localhost:50001
```

