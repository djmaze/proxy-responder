# Proxy responder

Listens for TCP connections on port 8080. Responds with the client's address and closes the connection.

Supports the [proxy protocol](https://www.haproxy.com/blog/haproxy/proxy-protocol/) for getting the real connection information when running behind a load balancer which supports the protocol as well (ie. HAProxy).

## Usage

Try it out with:

    docker run -d -p 8080:8080 decentralize/proxy-responder
    telnet localhost 8080
