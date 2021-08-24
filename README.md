# Chord
Implementation of the p2p chord protocol in Golang

# Paper
The implemenation is basd off the design/protocol brought up in https://pdos.csail.mit.edu/papers/ton:chord/paper-ton.pdf

# Generate Golang gRPC files

Note: This target assumes you have followed the [guide](https://grpc.io/docs/languages/go/quickstart/) to install grpc plugins and updated your PATH
 
```
$ make gen
```

# Run
To run

```
$ go run main.go
```