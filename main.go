package main

import (
	"github.com/Thedtk24/menv/cmd"
)

var version = "dev"

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
