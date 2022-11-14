package main

import (
	"flag"
	"fmt"
)

func main() {
	cmd := flag.String("input", "Hello World!!!", "Enter your name")
	flag.Parse()
	fmt.Println(*cmd)
}

// Then go mod init <module-name>
// create a binary -> go build main.go
// Run it -> ./main --input <give-any-input-here>
