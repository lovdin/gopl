package main

import (
    // "bufio"
    "io/ioutil"
    "fmt"
    "os"
    "strings"
)

// // v1
// func main() {
//     counts := make(map[string]int)
//     input := bufio.NewScanner(os.Stdin)
//     for input.Scan() {
//         counts[input.Text()]++
//     }
//     fmt.Println("------------------------------")
//     // note: ignoring potential errors from input
//     for line, n := range counts {
//         if n > 1 {
//             fmt.Printf("%d\t%s\n", n, line)
//         }
//     }
// }

// // v2
// func main() {
//     counts := make(map[string]int)
//     files := os.Args[1:]
//     if len(files) == 0 {
//         countLines(os.Stdin, counts)
//     } else {
//         for _, arg := range files {
//             f, err := os.Open(arg)
//             if err != nil {
//                 fmt.Fprintf(os.Stderr, "dup: %v\n", err)
//                 continue
//             }
//             countLines(f, counts)
//             f.Close()
//         }
//     }
//     fmt.Println("------------------------------")
//     for line, n := range counts {
//         if n > 1 {
//             fmt.Printf("%d\t%s\n", n, line)
//         }
//     }
// }

// func countLines(f *os.File, counts map[string]int) {
//     input := bufio.NewScanner(f)
//     // note: ignoring potential errors from input
//     for input.Scan() {
//         counts[input.Text()]++
//     }
// }

// v3
func main() {
    counts := make(map[string]int)
    for _, filename := range os.Args[1:] {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }
        for _, line := range strings.Split(string(data), "\n") {
            counts[line]++
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
