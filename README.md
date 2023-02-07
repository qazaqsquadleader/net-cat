# net-cat
### Task
This project consists on recreating the NetCat in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

1. NetCat, nc system command, is a command-line utility that reads and writes data across network connections using TCP or UDP. It is used for anything involving TCP, UDP, or UNIX-domain sockets, it is able to open TCP connections, send UDP packages, listen on arbitrary TCP and UDP ports and many more.1.Your project must work in a similar way that the original NetCat works, in other words, you must create a group chat. The project must have the following features :

2. TCP connection between server and multiple clients (relation of 1 to many).
3. A name requirement to the client.
4. Control connections quantity.
5. Clients must be able to send messages to the chat.
6. Do not broadcast EMPTY messages from a client.
7. Messages sent, must be identified by the time that was sent and the user name of who sent the message, example : [2020-01-20 15:48:41][client.name]:[client.message]
8. If a Client joins the chat, all the previous messages sent to the chat must be uploaded to the new Client.
9. If a Client connects to the server, the rest of the Clients must be informed by the server that the Client joined the group.
10. If a Client exits the chat, the rest of the Clients must be informed by the server that the Client left.
11. All Clients must receive the messages sent by other Clients.
12. If a Client leaves the chat, the rest of the Clients must not disconnect.
13. If there is no port specified, then set as default the port 8989. Otherwise, program must respond with usage message: [USAGE]: ./TCPChat $port


### How to go run:
first you need enter /cmd/ dir
`cd cmd`
