package main

import (
    "net/http"
    "log"
    "io/ioutil"
    "encoding/json"
    "fmt"
)

func watch()  {
    url := "http://127.0.0.1:8888/watch"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatalf("Fail to request %v: %v\n", url, err)
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    defer resp.Body.Close()
    if err != nil {
        log.Fatalf("Fail to do req: %v\n", err)
    }

    if resp.StatusCode != http.StatusOK {
        defer resp.Body.Close()
        log.Fatal("for request '%+v', got status: %v", url, resp.StatusCode)
    }

    for {
        bs := []byte{}

        dec := json.NewDecoder(resp.Body)
        if dec.Decode(bs) != nil {
            fmt.Printf("aaa")
        }

        //resp.Body.Read(bs)
        if len(bs) > 0 {
            log.Println(string(bs))
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

