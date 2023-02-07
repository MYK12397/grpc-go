# Simple gRPC-go server and client


## Features

### Unary API
- Client sends a request and server returns the response.
### Server Streaming
- Client sends a request to the server and gets a stream back, it's a sequence of messages.
### Client Streaming
- Client sends a stream of messages to the server and server just sends a response.
### Bi-directional Streaming
- Both client and server communicates using streams event though it's a "stream" not a "queue", the order of the messages is preserved.