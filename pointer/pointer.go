package main

import "fmt"

func main() {
    var a int = 1
    var b *int = &a
    var c **int = &b
    var x int = *b

    fmt.Println("a = ", a) // 1
    fmt.Println("&a = ", &a) // 0xaaaaaaaa
    fmt.Println("*&a = ", *&a) // 1
    fmt.Println("b = ", b) // 0xaaaaaaaa
    fmt.Println("&b = ", &b) // 0xbbbbbbbb
    fmt.Println("*&b = ", *&b) // 0xaaaaaaaa
    fmt.Println("*b = ", *b) // 1
    fmt.Println("c = ", c) // 0xbbbbbbbb
    fmt.Println("*c = ", *c) // 0xaaaaaaaa
    fmt.Println("&c = ", &c) // 0xcccccccc
    fmt.Println("*&c = ", *&c) // 0xbbbbbbbb
    fmt.Println("**c = ", **c) // 1
    fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c) // 1
    fmt.Println("x = ", x) // 1
}
