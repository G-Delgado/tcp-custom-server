package main

type commandID int

const (
	CMD_NICK commandID = iota
	// More for every command
)

type command struct {
	id     commandID
	client *client
	args   []string
}
