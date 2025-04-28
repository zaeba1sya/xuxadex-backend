package sockets

import (
	"errors"
	"fmt"
	"slices"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type socket struct {
	id   string
	conn *websocket.Conn
}

type SocketsPool struct {
	mx   sync.Mutex
	pool map[string][]*socket
}

func NewSocketsPool() *SocketsPool {
	return &SocketsPool{
		pool: make(map[string][]*socket),
	}
}

func (sp *SocketsPool) AddConnection(activityID string, newConn *websocket.Conn) string {
	if _, ok := sp.pool[activityID]; !ok {
		sp.pool[activityID] = make([]*socket, 0)
	}

	id := uuid.New().String()

	sp.pool[activityID] = append(sp.pool[activityID], &socket{
		id:   id,
		conn: newConn,
	})

	return id
}

func (sp *SocketsPool) SendToAll(activityID string, msg any) error {
	if _, ok := sp.pool[activityID]; !ok {
		return nil
	}

	failedCount := 0
	for _, socket := range sp.pool[activityID] {
		if err := socket.conn.WriteJSON(msg); err != nil {
			failedCount++
		}
	}

	if failedCount > 0 {
		return errors.New(fmt.Sprintf("failed sending msg to %d connections", failedCount))
	}

	return nil
}

func (sp *SocketsPool) IsActivityPoolExists(activityID string) bool {
	_, ok := sp.pool[activityID]
	return ok
}

func (sp *SocketsPool) IsActivityPoolEmpty(activityID string) bool {
	if _, ok := sp.pool[activityID]; !ok {
		return true
	}

	return len(sp.pool[activityID]) == 0
}

func (sp *SocketsPool) RemoveConnection(activityID string, connID string) {
	if _, ok := sp.pool[activityID]; !ok {
		return
	}
	isLastConection := false
	if len(sp.pool[activityID]) <= 1 {
		isLastConection = true
	}

	for i, socket := range sp.pool[activityID] {
		if socket.id == connID {
			sp.pool[activityID] = slices.Delete(sp.pool[activityID], i, i+1)
			return
		}
	}

	if isLastConection {
		delete(sp.pool, activityID)
	}
}
