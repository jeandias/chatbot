// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jeandias/chatbot/bot"
	watson "github.com/jeandias/chatbot/watson"
	"github.com/joho/godotenv"
)

var addr = flag.String("addr", GetPort(), "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func startChatHub(hub *bot.Hub) {
	go hub.Run()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print(err)
	}

	watson.StartWatsonAssistant()

	flag.Parse()
	chatbot := bot.NewAgent()
	hub := bot.NewHub(chatbot)
	startChatHub(hub)

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)
	r.Handle("/ws", bot.ServeWs(hub))

	err = http.ListenAndServe(*addr, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4000"
		log.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
