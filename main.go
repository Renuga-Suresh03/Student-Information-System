// main.go

package main

import "controllers/server"

func main() {
	s := server.NewServer()
	err := s.Start("8081")
	if err != nil {
		panic(err)
	}
}
