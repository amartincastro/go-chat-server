# go-chat-server
:phone: Chat server written in Go, JS, and Vue

Based on:
- https://scotch.io/bar-talk/build-a-realtime-chat-server-with-go-and-websockets

Compatibility:
- The application was tested and works on Mozilla Firefox v68.0.2
- Certain features may not be available on other browsers or versions

To run:
- Open bash or terminal and enter "git clone https://github.com/amartincastro/go-chat-server.git"
- Navigate to the "src" folder using by entering "cd src"
- Enter "go run main.go" to launch the application on Port 8000
- Open Firefox and navigate to localhost:8000

To edit port:
- Open the file src/main.go in a text editor
- Edit the first argument of the http.ListenAndServe function to reflect the port of your choice (the default is :8000)
