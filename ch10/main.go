package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader to allow connections from any origin
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin
	},
}

// Handle WebSocket connections
func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Read messages from the WebSocket connection and echo them back to the client
	for {
		messageType, p, err := conn.ReadMessage() // Read the incoming message
		if err != nil {
			log.Println(err)
			return
		}

		// Echo the received message
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// Serve static HTML file on the root URL path ("/")
func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "t1.html")  // Serve the HTML file
}

func main() {
	// Serve the HTML file at the root path
	http.HandleFunc("/", serveHTML)
	
	// Set up the WebSocket handler
	http.HandleFunc("/ws", handleConnections)

	// Start the web server on port 2380
	log.Println("Starting server on port 5380...")
	if err := http.ListenAndServe(":5380", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
