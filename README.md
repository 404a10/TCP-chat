# TCP-chat

This is a TCP chat **server.**
<br />
Clients can access it via telnet.
<br />
Since this is part of a school assignment, I'll be explaining every file in the project.

## This is a fork from [plutov's repo](https://github.com/plutov/packagemain/tree/master/20-tcp-chat/chat) so, if you are interested please check it out.

### The command object
It defines the value of the commands that the client will use to asign it later in the client object.

### The client object
Defines the parameters that the user can interact with, like the nickname, the room the client will be connecting to or the commands the user can send to the server.
In the readInput function, we "glue" the input command with it's function sending it through a channel to the server.
Also, the msg and err function are written here.

### The room object
It defines a room as the name of the room and the members connected there.
The broadcast function is also defined here.

### The server object
This file defines the server as the rooms on that server and the commands sent to it.
In here is written how the server should serve the clients and the functions of the commands are defined.
