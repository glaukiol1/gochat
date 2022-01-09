# gochat
Simple chat communication app over TCP, wrriten in Golang.

## How does it work?

Each client starts a TCP server, in a port that is defined by the second command-line argument (*os.Args[1]*). It waits for data to be sent, then parses it and it outputs to the Stdout. To send a message, there is a simple TCP client that sends the data to the TCP server and closes the connection.

## How to test it?

Run two client instances locally. Open two terminals, and run;

```sh
# terminal 1
$> go run client.go 1000 # where 1000 is the port number
```

```sh
# terminal 2
$> go run client.go 1001 # wheere 1001 is the port number
```

Then it would ask you in each terminal to whom to connect to, in terminal one write `localhost:1001` and in terminal two write `localhost:1000`

Type your message and press enter to send it, you should now see it in the other terminal.
