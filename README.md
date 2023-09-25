a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?
    No packages

b) Does your implementation use threads or processes? Why is it not realistic to use threads?
    We needed easy communication via channels, so gorounties make the most sense.

c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?
    We would handle message reordering by saving the sequence start number and reorder by finding the message that has the sequence number+1 for each message

d) In case messages can be delayed or lost, how does your implementation handle message loss?
    Message loss is handled by telling the server how many messages are coming and counting the messages received

e) Why is the 3-way handshake important?
    The 3-way handshake is important because it makes the connection more reliable. It is a way of controlling the transmission of data by first establishing a connection and then beginning data transmission in indexed segments, where the (likely) sequence start is communicated before the actual transmission of data