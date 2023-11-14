To start the system nodes must be connected manually one by one. To do this open a terminal for each node and write the following with distinct names and ports that connect the nodes in a ring

go run client/client.go -name <name> -clientp <port> -serverp <port>

One of the nodes must start the circulation of the token. Therefore another flag can be set

-hastoken true

To see the process of which node gets access and when, check out the log files generated.