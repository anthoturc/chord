# Chord
Implementation of the p2p chord algorithm in Golang.

This is a client library that is meant to be consumed by other applications.

# Paper
The implemenation is basd off the design/protocol brought up in https://pdos.csail.mit.edu/papers/ton:chord/paper-ton.pdf

# Generate Golang gRPC files


Note: This target assumes you have followed the [guide](https://grpc.io/docs/languages/go/quickstart/) to install grpc plugins and updated your PATH. We are using plugins for APIv2. See more here: https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go

Specifically you should have the newer protoc-gen-go and protoc-gen-grpc-go plugins installed. You can check these out at https://pkg.go.dev/

 
```
$ export PATH=$PATH:$(go env GOPATH)/bin # maybe it's already on your path (:
$ make gen
```

# Usage
```
// Start the Chord RPC Server
chordConfig := &ChordServerConfig{
	Create: true,
	Join: false,
	RemoteNodeIpAddr: "0.0.0.0:8080"
}

node := service.Init(chordCofig)

...

node.DumpInfo()

...

addr, err := node.Lookup(key)

...


s.GracefulStop()
```

# Run
```
# Create a chord ring with a single chord node
$ go run main.go -create

# Add a chord node to an existing ring
$ go run main.go -join -remote-addr <ip:port>
```

# Issues
The finger table optimization is not currently implemented. So this implementation effectively organizes the nodes into a circular doubly-linked list. The finger table optimization is on the roadmap.

# TODO:
If a node joins the chord ring it should reach out to its successor to get all of its data. This can be done by having the new node paginate through all of its successor's data (sorted by ID) until a given returned ID > n.ID.

If a node leaves the chord ring it should reach out to its successor to put all of its data.
This can be done by having the leaving node batch put all of its data onto the successor node.