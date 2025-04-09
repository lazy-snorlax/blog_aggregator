package main

import (
	"errors"
)

type command struct {
	name string
	args []string
}

type commands struct {
	list map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.list[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.list[cmd.name]
	if !ok {
		return errors.New("command not found in map")
	}
	return f(s, cmd)
}
