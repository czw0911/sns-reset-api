// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libs

//import "time"

type Msg struct {
	ID string
	Contents []byte
}


var HubWS = NewHub()


type Hub struct {
	// Registered connections.
	connections map[string]*ConnWS

	// Register requests from the connections.
	Register chan *ConnWS

	// Unregister requests from connections.
	Unregister chan *ConnWS
	
	Broadcast chan *Msg
}


func (h *Hub) Run() {
	//ticker := time.NewTicker(pingPeriod)
	for {
		select {
		case c := <-h.Register:
			h.connections[c.ID] = c
		case c := <-h.Unregister:
			if _, ok := h.connections[c.ID]; ok {
				delete(h.connections, c.ID)
				close(c.Send)
			}
		case m := <-h.Broadcast:
			if c, ok := h.connections[m.ID]; ok {
				 c.Send <- m.Contents
			}
		//case <-ticker.C:
		}
	}
}


func NewHub() *Hub {
	return &Hub{
		Register : make(chan *ConnWS),
		Unregister : make(chan *ConnWS),
		Broadcast : make(chan *Msg),
		connections : make(map[string]*ConnWS),
	}
}