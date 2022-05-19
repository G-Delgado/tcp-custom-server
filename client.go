package main

import "net"

type client struct {
	conn     net.Conn
	nick     string
	channel  *channel
	commands chan<- command
}
