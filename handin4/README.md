To start the system nodes must be connected manually one by one. To do this open a terminal for each node and write the following with distinct names and ports that connect the nodes in a ring

go run client/client.go -name <name> -port <port> -nextport <port of next node>

Then to request access to the critical section write anything

ex. "ACCESS REQUEST"

To see the process of which node gets access and when, checl out the log files generated.