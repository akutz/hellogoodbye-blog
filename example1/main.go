package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {
	l, err := net.Listen("unix", "http.sock")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	http.Serve(l, http.FileServer(http.Dir(".")))
}
