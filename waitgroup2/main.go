package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func printName(name string, wg *sync.WaitGroup) {
    ms := rand.Intn(200) + 100
    time.Sleep(time.Duration(ms) * time.Millisecond)
    fmt.Println(name)
    wg.Done()
}

func main() {
    fmt.Println("********")

    var wg sync.WaitGroup
    names := []string{
        "Alice", "Bob", "Carl",
    }
    ch := make(chan string, 3)
    for _, name := range names {
        wg.Add(1)
        ch <- name
    }

    go func() {
        wg.Wait()
        close(ch)
    }()
    for name := range ch {
        go printName(name, &wg)
    }

    fmt.Println("********")
}