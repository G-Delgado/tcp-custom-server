package main

import "net"

type channel struct {
	name    string
	members map[net.Addr]*client
}

func (c *channel) broadcast(sender *client, msg string) {
	for addr, m := range c.members {
		if addr != sender.conn.RemoteAddr() {
			m.msg(msg)
		}
	}
}
