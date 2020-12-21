// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bot

import "encoding/json"

type ClientMessage struct {
	client  *Client
	message []byte
}

type SocketMessage struct {
	Event   string
	Message string
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	chatbot Bot
	// Registered clients.
	clients map[*Client]bool

	broadcastmsg chan *ClientMessage
	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub(chatbot Bot) *Hub {
	return &Hub{
		chatbot:      chatbot,
		broadcastmsg: make(chan *ClientMessage),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		clients:      make(map[*Client]bool),
	}
}

func (h *Hub) SendMessage(client *Client, message []byte) {
	client.send <- message
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			greeting := h.chatbot.Greeting()
			h.SendMessage(client, []byte(greeting))
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case clientmsg := <-h.broadcastmsg:
			client := clientmsg.client

			var clientJson SocketMessage
			json.Unmarshal([]byte(clientmsg.message), &clientJson)

			reply := h.chatbot.Reply(clientJson.Event, clientJson.Message)
			h.SendMessage(client, []byte(reply))
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
