## TCP tls server/client

### Description
This project is a simple TCP + tls server/client implementation

Client sends a message to the server and 
the server responds with the message about bytes received

Need run the server first and then the client
If client is disconnected, the server can connect to another client

### How to run
1. Run the TCP server
```./hw8 -mode=tcp-server -ip=[ip] -key=[key_file] -crt=[certificate_file]```
2. Run the TCP client
```./hw8 -mode=tcp-client -ip=[ip] -key=[key_file] -crt=[certificate_file]```