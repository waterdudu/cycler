package connection

import (
	uuid "github.com/satori/go.uuid"
	"github.com/waterdudu/cycler/hub"
)

type Connection struct {
	ID      uuid.UUID
	shortId string
	hub     *hub.Hub
}
