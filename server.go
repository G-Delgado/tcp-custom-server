package main

type server struct {
	channels map[string]*channel
	commands chan command
}
