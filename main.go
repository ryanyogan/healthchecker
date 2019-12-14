package main

import (
	"flag"
	"net/http"
	"os"
)

// Main starts a very simple application that pings defines ports based
// on the command line args (flag)
func main() {
	port := flag.String("port", "80", "port on localhost to check")
	flag.Parse()

	resp, err := http.Get("http://127.0.0.1:" + *port + "/health")

	if err != nil || resp.StatusCode != 200 {
		os.Exit(1)
	}
	os.Exit(0)
}
