package main

import (
    "net/http"
    "log"
    "io/ioutil"
    "testing/iotest"
    "io"
)

func watch()  {
    url := "http://127.0.0.1:8888/watch"
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalf("Fail to GET %v: %v\n", url, err)
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        log.Fatal("for request '%+v', got status: %v", url, resp.StatusCode)
    }

    reader := iotest.OneByteReader(resp.Body)
    buf := make([]byte, 0, 1024)
    onebyte := make([]byte, 1)
    for {
        _, err := reader.Read(onebyte)

        if err == io.EOF {
            if len(buf) > 0 {
                log.Println(string(buf))
            }
            log.Printf("disconnect to %v\n", resp.Request.RemoteAddr)
            break
        }

        switch onebyte[0] {
        case '\n':
            log.Println(string(buf))
            buf = buf[0:0]
        default:
            buf = append(buf, onebyte[0])
        }
    }
}

func list()  {
    url := "http://127.0.0.1:8888/watch"
    resp, err := http.Get(url)

    if err != nil {
        log.Fatalf("Fail to get %v: %v\n", url, err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Fail to read body: %v\n", err)
    }
    log.Println(string(body))
}

func main() {
    //list()
    watch()
}
