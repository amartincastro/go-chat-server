package main

// Todo list: 
// Timestamp of messages sent to server and logs
// Timestamp of messages sent to other users
// 

import (
	"log"
	"net/http"
	"time"
	"io/ioutil"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)		// Connected clients. The first variable is a map where the key is actually a pointer to a WebSocket. This data structure contains all connected clients to the server.

var broadcast = make(chan Message)         			 // Croadcast channel. This variable is a channel that will act as a queue for messages sent by clients. 

var upgrader = websocket.Upgrader{} 				 // Configure the upgrader - an object with methods for taking a normal HTTP connection and upgrading it to a WebSocket as we'll see later in the code.

type Message struct {
	Email    string `json:"email"`					 // Define our message object
	Username string `json:"username"`
	Message  string `json:"message"`
	Timestamp int time.Now()
}




func main() {
    fs := http.FileServer(http.Dir("../public"))	// Create a file server to route web visitors to our JS app
    http.Handle("/", fs)

	http.HandleFunc("/ws", handleConnections)		// Configure websocket route to handle any requests for initializing a websocket, and pass in handleConnections function as argument

	go handleMessages() 					// Initialize goroutine for incoming chat messages

	log.Println("http server started on :8000")		// Start the server on localhost port 8000 and log any errors. ---- Add function here to be able to modify port # ----
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	func handleConnections(w http.ResponseWriter, r *http.Request) { // Upgrade initial GET request from cleints into websockets. 
        ws, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
                log.Fatal(err)
        }

		defer ws.Close()							// Make sure we close the connection when the function returns
		
		clients[ws] = true							// Register new clients to the map "clients"
		
		for {										// Infinite loop that continuously waits for a new message to be written to the WebSocket

			var msg Message
			err := ws.ReadJSON(&msg)				// Read in a new message as JSON and map it to a Message object
			if err != nil {
					log.Printf("error: %v", err)
					delete(clients, ws)				// If there is an error from the WebSocket, log the error and remove that client from our global "clients" map so we don't try to read from or send new messages to that client
					break
			}
			broadcast <- msg						// Send the new message to the broadcast channel
		}
	}

	func handleMessages() {
        	for {										 // Grab the next message from the broadcast channel
                msg := <- broadcast
                for client := range clients {		 // Send msg to connected clients

                        err := client.WriteJSON(msg)
                        if err != nil {
                                log.Printf("error: %v", err)
                                client.Close()
                                delete(clients, client)
                        }
                }
        	}
	}

	func (m Message) saveToFile(filename string) error {
		return ioutil.WriteFile(filename, []byte(m.toString()), 0666)
	}

	Message.saveToFile("chat_logs") // Save messages to local log file

}
