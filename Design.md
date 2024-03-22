# Project Design Overview

## Project Description

This is a neurotech headset simulator built in Golang. It serves as a simulation for what a neurotech headset's output is supposed to look like. The simulator generates random EEG data and sends it to a server. The server then processes the data and sends it to a client. The client is disjoint and can be found as one of the other repositories in this organization.

Note that the logic here also serves as a reference for bulding other server side applications for passing neural data from the headset to the client.

## Design Decisions

### Key Design Considerations

This project needs to send out a rapid amount of data every second. This simulation is designed with the understanding that there might be cases where over $10000$ packets are sent out every second. This is why the server is designed to be able to handle a large amount of data. It is critical that this server is able to handle a large amount of data.

Another thing to factor into the design is that the frontend should be able to easily work with and use the data in an easy to read way. It's pivotal to reduce any needed client work and make the server do as much of the work as possible.

### Design Options

There are a few different options for what could be done. We have:

1. WebSockets
2. Server Sent Events (SSE)
3. gRPC

> WebSockets

Let's break each down to see which options would work best and why. Starting off, we have **WebSockets**. WebSockets enable us to open interactive sessions between users and servers. They're a standard solution for building real time applications. WebSockets are a full-duplex communication channel over a single Transmission Control Protocol (TCP) connection enabling the client and server to send data without requiring any requests to be made. This is good option for anything that requires continuous data exchange. They require the remote host to opt in to that code and the main method of security for this method is the origin-based security modeal commonly used in web browsers. All the information for WebSockets is present inn the [`RFC 6455 specification`](https://datatracker.ietf.org/doc/html/rfc6455).

WebSockets have two key parts to them, the handshake and the data transfer. The handshake is requested by the browsers and responded by the servers, after whic, a connection is established. WebSockets have a header that only requires a handshake between a browser and server for establishing a connection. Once the client and server have both sent their handshake, and if the handshake was successful, then the data exchange part continues.

Golang has a few libraries for WebSockets including [`gorilla/websocket`](https://github.com/gorilla/websocket) and [`net/http`](https://pkg.go.dev/net/http).

> Server Sent Events (SSE)

**SSE** are a standard describing how servers can initiate data transmission towards the client once an initial client connection has been established. They are commonly used to send real-time updates to a web application. They are a one-way channel from the server to the client

> gRPC
