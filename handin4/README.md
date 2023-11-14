To start the system nodes must be connected manually one by one. To do this open a terminal for each node and write the following with distinct names and ports that connect the nodes in a ring

go run client/client.go -name <name> -clientp <port> -serverp <port>

One of the nodes must start the circulation of the token. Therefore another flag can be set

-hastoken true

example set up:

terminal 1
> go run node.go -name hanna -clientp 5400 -serverp 5401

terminal 2
> go run node.go -name oskar -clientp 5401 -serverp 5402

terminal 3
> go run node.go -name jeppe -clientp 5402 -serverp 5400 -hastoken true

To see the process of which node gets access and when, check out the log files generated.