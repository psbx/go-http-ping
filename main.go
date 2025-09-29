package main

import (
    "fmt"
    "net/http"
    "os"
    "sync"
    "time"
)

func ping(url string, wg *sync.WaitGroup) {
    defer wg.Done()
    t0 := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("%s ERROR: %v\n", url, err)
        return
    }
    resp.Body.Close()
    fmt.Printf("%s %d %v\n", url, resp.StatusCode, time.Since(t0))
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("usage: go-http-ping <url1> <url2> ...")
        os.Exit(1)
    }
    var wg sync.WaitGroup
    for _, u := range os.Args[1:] {
        wg.Add(1)
        go ping(u, &wg)
    }
    wg.Wait()
}
