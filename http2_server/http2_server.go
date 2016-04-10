package main

import (
    "github.com/bradfitz/http2"
    "net/http"
    "log"
    "time"
    "io"
    "strings"
)

func main() {
    srv := http.Server{Addr:":8888"}

    registerHandlers()

    log.Println("Listen on " + srv.Addr)

    http2.ConfigureServer(&srv, &http2.Server{})

    go func() {
        log.Fatal(srv.ListenAndServe())
    } ()

    select { }
}

func registerHandlers()  {
    mux := http.NewServeMux()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        mux.ServeHTTP(w, r)
    })
    
    mux.HandleFunc("/watch/", watchHandler)
}

func watchHandler(w http.ResponseWriter, r *http.Request) {
    clientGone := w.(http.CloseNotifier).CloseNotify()
    w.Header().Set("Content-Type", "text/plain")
    defer func() {
        if rv := recover(); rv != nil {
            log.Fatal(rv)
        }
    } ()

    io.WriteString(w, "# ~1KB of junk to force browsers to start rendering immediately: \n")
    io.WriteString(w, strings.Repeat("# xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n", 13))

    for {

        select {
            case <-clientGone :
                log.Printf("Client %v disconnected from the watch", r.RemoteAddr)
                break
            default:
                io.WriteString(w, time.Now().String() + "\n")

                w.(http.Flusher).Flush()
                time.Sleep(time.Second)
        }
    }
}