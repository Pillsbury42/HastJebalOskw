The system only works with specific ports, to make the connection between clients and servers simpler. To start the program, run the following commands:

go run rm/server.go
go run rm/server.go -port 5401
go run rm/server.go -port 5402
go run client/client.go -name <name>
go run client/client.go -name <name>

Each command should be run in its own terminal. The server commands should be run before the client commands, other than that the order is not important. You can create any number of clients, though at least two are necessary to explore the full functionality of the program. You cannot create more servers without directly changing the code to support so.

In the client terminals, you can interact with the auction by writing the following two commands:

result
bid <integer>

When the clients start up, a "help" message is printed which also explains this. To see what happens in the program when it's running (when elections happen, for instance), consult the log files that are created.