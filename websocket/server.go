package main

import (
    "code.google.com/p/go.net/websocket"
    "fmt"
    "html/template"
    "log"
    "net/http"
)

func Echo(ws *websocket.Conn) {
    var err error

    for {
        var reply string

        if err = websocket.Message.Receive(ws, &reply); err != nil {
            fmt.Println("Can't receive")
            break
        }

        fmt.Println("Received back from client: " + reply)

        msg := "Received:  " + reply
        fmt.Println("Sending to client: " + msg)

        if err = websocket.Message.Send(ws, msg); err != nil {
            fmt.Println("Can't send")
            break
        }
    }
}

func loadform(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        fmt.Println("Receive loadform request");
        t, _ := template.ParseFiles("client.html")
        t.Execute(w, nil)
    }
}

func main() {
    http.HandleFunc("/loadform", loadform)
    http.Handle("/gowebsocket", websocket.Handler(Echo))

    if err := http.ListenAndServe(":1234", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
