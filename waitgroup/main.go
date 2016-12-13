package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func printName(ch <-chan string, wg *sync.WaitGroup) {
    for name := range ch {
        ms := rand.Intn(200) + 100
        time.Sleep(time.Duration(ms) * time.Millisecond)
        fmt.Println(name)
        wg.Done()
    }
}

func main() {
    fmt.Println("********")

    names := []string{
        "Alice", "Bob", "Carl",
    }
    ch := make(chan string, 3)
    for _, name := range names {
        ch <- name
    }

    var wg sync.WaitGroup
    wg.Add(len(names))
    go printName(ch, &wg)

    wg.Wait()
    close(ch)

    fmt.Println("********")
}
