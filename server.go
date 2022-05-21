package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	channels map[string]*channel
	commands chan command
}

func newServer() *server {
	return &server{
		channels: make(map[string]*channel),
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_CHANNELS:
			s.listchannels(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.args)
		}
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("New client has connected: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		nick:     "anonymous",
		commands: s.commands,
	}

	c.readInput()

}

func (s *server) nick(c *client, args []string) {
	c.nick = args[1]
	c.msg(fmt.Sprintf("All  right, I will call  you %s", c.nick))
}

func (s *server) join(c *client, args []string) {
	channelName := args[1]
	r, ok := s.channels[channelName]
	if !ok {
		r = &channel{
			name:    channelName,
			members: make(map[net.Addr]*client),
		}
		s.channels[channelName] = r
	}
	s.quitCurrentchannel(c)
	r.members[c.conn.RemoteAddr()] = c
	r.broadcast(c, fmt.Sprintf("%s has joined  the channel", c.nick))
	c.msg(fmt.Sprintf("Welcome to %s", r.name))

	if c.channel != nil {

	}
	c.channel = r
}
func (s *server) listchannels(c *client, args []string) {
	var channels []string
	for name := range s.channels {
		channels = append(channels, name)
	}

	c.msg(fmt.Sprintf("Availabe channels are : %s", strings.Join(channels, ", ")))
}

func (s *server) msg(c *client, args []string) {
	if c.channel == nil {
		c.err(errors.New("You must join the channel first"))
		return
	}

	c.channel.broadcast(c, c.nick+": "+strings.Join(args[1:], " "))
}

func (s *server) quit(c *client, args []string) {
	log.Printf("Client  has disconnected: %s", c.conn.RemoteAddr().String())
	s.quitCurrentchannel(c)

	c.msg("Sad to see you go :(")

	c.conn.Close()
}

func (s *server) quitCurrentchannel(c *client) {
	if c.channel != nil {
		delete(c.channel.members, c.conn.RemoteAddr())
		c.channel.broadcast(c, fmt.Sprintf("%s has left the channel", c.nick))
	}
}
