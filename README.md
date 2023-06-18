# Distributed Go

A fairly simple example of a distributed system written entirely in Go.

**NOTE:** Don't take this as a "best practice" for how distributes systems should be architected. This is more of a learning resource for the concepts. e.g. there's no containers for coordinated builds etc.

For production ready distributed services tools, check out [Go-kit](https://gokit.io).

For production ready web services, check out [Go Micro](https://go-micro.dev).

## Start it up

Build the services.
```shell
go build ./cmd/registryservice
go build ./cmd/logservice
go build ./cmd/gradingservice
go build ./cmd/teacherportalservice
```

Open 4 terminals. Run the following 4 commands, one in each terminal, making sure to start the `registryservice` first.

```shell
./registryservice
./logservice
./gradingservice
./teacherportalservice
```

Notice that as the `logservice`, `gradingservice` and `teacherportalservice` start, they register with the `registryservice`, and dependant services become aware of their dependencies starting.

Also note the heartbeat checks for service monitoring.

## About this project

The type is **Hybrid**, with a central **hub** service as a service registry with health monitoring. The rest of the services follow a peer-to-peer model, including logging and any business logic.

The language is Go, with no frameworks, and communication is (initially) over HTTP with JSON as the protocol.

### Main components

**Service registry**
- Service registration and de-registration.
- Health monitoring.

**Teacher portal**
- Web application.
- API Gateway.

**Log service**
- Centralized logging.

**Grading service**
- Business logic.
- Data persistence (no database).

## Types of distributed systems

### Hub and Spoke

**Centralized service** (hub) coordinates other services, distributing requests to them and handling responses.

**Pros**
* The hub can act as a load balancer and a gateway.
* It's also good for tracing and logging.
* Simpler service discovery.

**Cons**
* Single point of failure.
* Multiple roles with increased complexity (not all that bad).

### Peer to peer

Each service can communicate with one another, and have the ability to let other services know when it starts up and shuts down.

**Pros**
* No single point of failure, quire fault tolerant.
* Highly decoupled.

**Cons**
* Tricky service discovery.
* Tricky load balancing.

### Message queues

Services send messages to queues, and other services can subscribe to those queues, as well as send to queues themselves. Known as event driven architecture.

**Pros**
* Easy to scale.
* Very decoupled.
* Message persistence (usually built in to queue provider), making it more fault tolerant.

**Cons**
* Queues (depending on arch) can be the single point of failure.
* Updating configuration on messages/types can be tricky with so many moving parts.

### Hybrid (most common)

As suggested, a mixture of the others. Often there's a **hub** that handles _some_ routing to other services, but can also act as another peer.

**Pros**
* Can have simple load balancing.
* Quite robust against failure.

**Cons**
* Can easily become overly complex and difficult to understanding.
* Hub can easily gain scope creep.

## Things to consider when architecting distributed systems

### Languages & Frameworks
Single or multiple languages? Are the frameworks compatible or do they need an additional transform layer?

Does the lang/framework support the transport layers and protocols?

Also think about stability and ecosystem.

### Transports
How are the services going to interact and communicate. TCP/UDP/RPC/HTTP or a mix? Note, mix can provide a lot of overhead. i.e setting up support so the same requests can be processed on multiple transport types.

### Protocols
The types of messages that will be sent. These could be language specific (think gob in go), or JSON, or Protocol Buffers for binary? Efficiency and complexity tradeoffs apply.