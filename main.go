package main

import (
	"fmt"
	"os"

	"github.com/datamammoth/packer-plugin-datamammoth/builder/datamammoth"
)

// version is set at build time via ldflags.
var version = "0.1.0-dev"

func main() {
	pps := datamammoth.PluginSet()
	if pps == nil {
		fmt.Fprintln(os.Stderr, "packer-plugin-datamammoth: plugin set is nil")
		os.Exit(1)
	}

	// In production, this calls pps.Run() which implements
	// the Packer plugin protocol. For now we print version info.
	fmt.Printf("packer-plugin-datamammoth %s\n", version)
	fmt.Println("Use with: packer build -only=datamammoth.server ...")
}
