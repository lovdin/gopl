package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "time"
    "sync"
)

var done = make(chan struct{})

func cancelled() bool {
    select {
    case <-done:
        return true
    default:
        return false
    }
}

var vFlag = flag.Bool("v", false, "show verbose progress messages")

func main() {
    flag.Parse()
    roots := flag.Args()
    if len(roots) == 0 {
        roots = []string{"."}
    }

    go func () {
        os.Stdin.Read(make([]byte, 1))
        close(done)
    }()

    fileSizes := make(chan int64)
    var wg sync.WaitGroup
    for _, root := range roots {
        wg.Add(1)
        go walkDir(root, &wg, fileSizes)
    }
    go func () {
        wg.Wait()
        close(fileSizes)
    }()

    var tick <-chan time.Time
    if *vFlag {
        tick = time.Tick(time.Second)
    }
    var nfiles, nbytes int64
loop:
    for {
        select {
        case <-done:
            for range fileSizes {
                // Do nothing. Just drain fileSizes to allow existing goroutines to finish.
            }
            return
        case size, ok := <-fileSizes:
            if !ok {
                break loop
            }
            nfiles++
            nbytes += size
        case <- tick:
            printDiskUsage(nfiles, nbytes)
        }
    }
    printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
    fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
    defer wg.Done()

    if cancelled() {
        return
    }

    for _, entry := range dirents(dir) {
        if entry.IsDir() {
            wg.Add(1)
            subdir := filepath.Join(dir, entry.Name())
            go walkDir(subdir, wg, fileSizes)
        } else {
            fileSizes <- entry.Size()
        }
    }
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
    select {
    case sema <- struct{}{}:
        // Do nothing. Just acquire token
    case <-done:
        return nil // Cancelled
    }
    defer func () {
        <-sema // Release token
    }()

    entries, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Fprintf(os.Stderr, "customdu: %v\n", err)
        return nil
    }
    return entries
}
