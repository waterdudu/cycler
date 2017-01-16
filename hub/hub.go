package hub

import (
	"github.com/waterdudu/cycler/connection"
)

type Hub struct {
	connections map[*connection.Connection]bool

	register   chan *connection.Connection
	unregister chan *connection.Connection
}

// NewHub initializes a new hub.
func NewHub() (*Hub, error) {
	return &Hub{
		connections: make(map[*connection.Connection]bool),
		register:    make(chan *connection.Connection),
		unregister:  make(chan *connection.Connection),
	}, nil
}
