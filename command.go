package main

type commandID int

const (
	CMD_NICK commandID = iota
	CMD_JOIN
	CMD_CHANNELS
	CMD_MSG
	CMD_QUIT
	// More for every command
)

type command struct {
	id     commandID
	client *client
	args   []string
}
