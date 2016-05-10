package main

import (
	"os"

	"github.com/intelsdi-x/snap/control/plugin"
	session "github.com/intelsdi-x/swan/misc/snap-plugin-publisher-session-test/session-publisher"
)

func main() {
	meta := session.Meta()
	plugin.Start(meta, new(session.SessionPublisher), os.Args[1])
}