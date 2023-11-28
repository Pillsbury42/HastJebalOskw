The system only works with specific ports, to make the connection between clients and servers simpler. To start the program, run the following commands:

go run rm/server.go
go run rm/server.go -port 5401
go run rm/server.go -port 5402
go run client/client.go -name oskar
go run client/client.go -name november -serverp 5401

The commands should be run in order, and each should run in its own terminal. In the client terminals, you can interact with the auction by writing the two following commands:

result
bid <integer>

When the clients start up, a "help" message is printed which also explains this. To see what happens in the program when it's running (when elections happen, for instance), consult the log files that are created.