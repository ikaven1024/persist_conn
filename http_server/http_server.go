package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	id = "server1"
)

func main() {
	log.Println(os.Args)
	if len(os.Args) > 1 {
		id = os.Args[1]
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/watch/", watchHandler)
	mux.HandleFunc("/list/", listHandler)

	server := &http.Server{
		Addr:    ":8888",
		Handler: mux,
	}
    log.Println("Listen on " + server.Addr)
	log.Fatal(server.ListenAndServe())
}

func watchHandler(w http.ResponseWriter, r *http.Request) {
	clientGone := w.(http.CloseNotifier).CloseNotify()
	w.Header().Set("Content-Type", "text/plain")
	defer func() {
		if rv := recover(); rv != nil {
			log.Fatal(rv)
		}
	}()

	io.WriteString(w, "# ~1KB of junk to force browsers to start rendering immediately: \n")
	io.WriteString(w, strings.Repeat("# xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n", 13))
	io.WriteString(w, "i am "+id+"\n")
	for {
		select {
		case <-clientGone:
			log.Printf("Client %v disconnected from the watch", r.RemoteAddr)
			break
		default:
			io.WriteString(w, "Watch: " + time.Now().String()+"\n")

			w.(http.Flusher).Flush()
			time.Sleep(time.Second)
		}
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "i am "+id+"\n")
    io.WriteString(w, "List: " + time.Now().String() + "\n")
}