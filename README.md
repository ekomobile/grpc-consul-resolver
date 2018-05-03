[![Build Status](https://travis-ci.org/ekomobile/grpc-consul-resolver.svg)](https://travis-ci.org/ekomobile/grpc-consul-resolver)
[![GitHub release](https://img.shields.io/github/release/ekomobile/grpc-consul-resolver.svg)](https://github.com/ekomobile/grpc-consul-resolver/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/ekomobile/grpc-consul-resolver)](https://goreportcard.com/report/github.com/ekomobile/grpc-consul-resolver)
![Downloads](https://img.shields.io/github/downloads/ekomobile/grpc-consul-resolver/total.svg)
[![GoDoc](https://godoc.org/github.com/ekomobile/grpc-consul-resolver?status.svg)](https://godoc.org/github.com/ekomobile/grpc-consul-resolver)

# gRPC Consul resolver
This lib resolves Consul services by name. 

# Usage

Somewhere in your `init` code: 
```go
import (
    "github.com/ekomobile/grpc-consul-resolver"
)

// Will query consul every 5 seconds.
resolver.RegisterDefault(time.Second * 5)
```

Getting connection:
```go
conn, err := grpc.DialContext(ctx, "srv://consul/my-awesome-service")
```
With round-robin balancer:
```go
import (
    "google.golang.org/grpc/balancer/roundrobin"
)

conn, err := grpc.DialContext(ctx, "srv://consul/my-awesome-service", grpc.WithBalancerName(roundrobin.Name))
```
