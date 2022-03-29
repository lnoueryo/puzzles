// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	_ "embed"
	"flag"
	cmd "backend/commands"
)

func main() {
	flag.Parse()
	cmd.CMD()
}


