# Chord
Implementation of the p2p chord protocol in Golang

# Paper
The implemenation is basd off the design/protocol brought up in https://pdos.csail.mit.edu/papers/ton:chord/paper-ton.pdf

# Generate Golang gRPC files

Note: This target assumes you have followed the [guide](https://grpc.io/docs/languages/go/quickstart/) to install grpc plugins and updated your PATH
 
```
$ make gen
```

# Usage
```
// Start the Chord RPC Server
chordConfig := &ChordServerConfig{
	Create: 					true,
	Join: 						false,
	RemoteNodeIpAddr: "0.0.0.0:8080"
}

chord := chord.Init(chordCofig)

...

chord.DumpInfo()

...

addr, err := chord.Lookup(key)

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