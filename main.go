package main

import (
	"flag"
	"fmt"

	"github.com/anthoturc/chord/service"
)

func main() {
	// Require options to create or join a chord ring
	create := flag.Bool("create", false, "Create a chord ring")
	join := flag.Bool("join", false, "Join an existing chord ring")
	remoteNodeIpAddr := flag.String("remote-addr", "", "The address of any node on an existing Chord ring. This is ignored if the create flag is specified")
	flag.Parse()

	chord := service.Init(
		&service.ChordServerConfig{
			Create:           *create,
			Join:             *join,
			RemoteNodeIpAddr: *remoteNodeIpAddr,
		},
	)

	for {
		fmt.Println("Enter option: l (lookup key), d (dump node info), q (quit)")
		fmt.Print("> ")
		var option string
		fmt.Scanf("%s", &option)

		switch option {
		case "d":
			fmt.Println(chord.DumpInfo())
		case "l":
			fmt.Print("Enter key: ")
			var key string
			fmt.Scanf("%s", &key)
			addr, _ := chord.Lookup(key)
			fmt.Println(addr)
		}

		if option == "q" {
			break
		}
	}

	chord.Stop()
}
