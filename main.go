package main

import (
	"fmt"
	"io"
	"net"

	// "net_test/utils"
	"os"
)

func main() {
	server, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Print("Sape")
	}
	conn, err := server.Accept()
	fi, err := os.Open("Sape")
	io.Copy(conn, fi)
	// utils.CheckError(err)

}
