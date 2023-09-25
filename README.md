a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?
    We use a struct called packet to hold the sequence number, the message length and a piece of the message. 

b) Does your implementation use threads or processes? Why is it not realistic to use threads?
    We needed easy communication via channels, so goroutines made the most sense. Threads run locally, therefore the risk of data being lost is very, very small, unlike on an actual network.

c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?
    The sequence number from the packet struct can be used to check what order the data should be in. The received message packets could be arranged in a slice and then sorted by their sequence number.

d) In case messages can be delayed or lost, how does your implementation handle message loss?
    The client waits for a maximum of 3 seconds for an answer from the server during the handshake. If it receives none, it restarts the routine
    Similarly, the server waits for 1 second for the acknowledgement from the client, and restarts otherwise

e) Why is the 3-way handshake important?
    The 3-way handshake is important because it makes the connection more reliable. 
    The handshakes ensures that data can be sent both ways, from the client to the server and vice versa, before you start to send the actual messages/data you want to transfer.