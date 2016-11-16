package main

import (
    "fmt"
    "os"
    "flag"
    "strings"
)

// // v1
// func main() {
//     var s, sep string
//     for i := 0; i < len(os.Args); i++ {
//         s += sep + os.Args[i]
//         sep = " "
//     }
//     fmt.Println(s)
// }

// // v2
// func main() {
//     s, sep := "", ""
//     for _, arg := range os.Args[0:] {
//         s += sep + arg
//         sep = " "
//     }
//     fmt.Println(s)
// }

// v3
func main() {
    flag.Parse()
    fmt.Println(strings.Join(flag.Args(), " "))
    fmt.Println(strings.Join(os.Args[1:], " "))
}
