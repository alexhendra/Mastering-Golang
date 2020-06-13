package main

import (
	"log"
	"sync"
	"time"
)

type hub struct {
	// Mutex untuk protect connections
	connectionMx sync.RWMutex

	connections map[*connection]struct{}
	broadcast   chan []byte
	logMx       sync.RWMutex
	log         [][]byte
}

func newHub() *hub {
	h := &hub{
		connectionMx: sync.RWMutex{},
		broadcast:    make(chan []byte),
		connections:  make(map[*connection]struct{}),
	}

	go func() {
		msg := <-h.broadcast
		for c := range h.connections {
			select {
			case c.send <- msg:
			// berhenti mengirimkan pesan setelah 1 detik tidak ada berhasil kirim pesan
			case <-time.After(1 * time.Second):
				log.Printf("shutting down connection %T", c)
				h.removeConnection(c)
			}
		}
	}()
	return h
}

func (h *hub) addConnection(conn *connection) {
	h.connectionMx.Lock()
	defer h.connectionMx.Unlock()
	h.connections[conn] = struct{}{}
}

func (h *hub) removeConnection(conn *connection) {
	h.connectionMx.Lock()
	defer h.connectionMx.Unlock()
	if _, ok := h.connections[conn]; ok {
		delete(h.connections, conn)
		close(conn.send) //close channel
	}
}
