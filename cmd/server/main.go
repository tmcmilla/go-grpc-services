package main

import "log"

func Run() error {
	// responsible for initiating and starting
	// our gRPC server
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
