package main

import (
	"flag"
	"fmt"

	"github.com/anthoturc/chord/chord"
)

func main() {
	// Require options to create or join a chord ring
	create := flag.Bool("create", false, "Create a chord ring")
	join := flag.Bool("join", false, "Join an existing chord ring")
	remoteNodeIpAddr := flag.String("remote-addr", "", "The address of any node on an existing Chord ring. This is ignored if the create flag is specified")
	flag.Parse()

	node, s := chord.Init(
		&chord.ChordServerConfig{
			Create:           *create,
			Join:             *join,
			RemoteNodeIpAddr: *remoteNodeIpAddr,
		},
	)

	for {
		fmt.Println("Enter option: d (dump node info), q (quit)")
		fmt.Print("> ")
		var option string
		fmt.Scanf("%s", &option)

		switch option {
		case "d":
			fmt.Println(node.DumpNodeInfo())
		}

		if option == "q" {
			break
		}
	}

	s.GracefulStop()
}
