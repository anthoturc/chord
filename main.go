package main

import (
	"flag"
	"log"

	"github.com/anthoturc/chord/chord"
)

func main() {
	// Require options to create or join a chord ring
	create := flag.Bool("create", false, "Create a chord ring")
	join := flag.Bool("join", false, "Join an existing chord ring")
	remoteAddrToJoin := flag.String("remote-addr", "", "The address of any node on an existing Chord ring. This is ignored if the create flag is specified")
	flag.Parse()

	chord.Init()

	if *create == *join {
		log.Fatal("Cannot create a chord ring and join one too")
	} else if *create {
		// on create set the successor to itself.
	} else {
		// on join, make remote find predecessor call
	}
}
