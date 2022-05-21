package main

import (
	"log"
	"net"
)

// "net_test/utils"

func main() {
	s := newServer()
	go s.run()
	// server, err := net.Listen("tcp", "localhost:8080")
	// if err != nil {
	// 	fmt.Print("Sape")
	// }
	// conn, err := server.Accept()
	// fi, err := os.Open("Sape")
	// io.Copy(conn, fi)
	// // utils.CheckError(err)

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Unable to start server: %s", err.Error())
	}
	defer listener.Close()
	log.Printf("Started server on :8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
