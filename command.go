package main

type commandID int

const (
	CMD_NICK commandID = iota
	CMD_SUBSCRIBE
	CMD_CHANNELS
	CMD_MSG
	CMD_QUIT
	CMD_SEND
	// More for every command
)

type command struct {
	id     commandID
	client *client
	args   []string
}
