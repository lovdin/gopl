package equal

import (
    "bytes"
    "fmt"
    "testing"
)

func TestEqual(t *testing.T) {
    one, oneAgain, two := 1, 1, 2

    type CyclePtr *CyclePtr
    var cp1, cp2 CyclePtr
    cp1 = &cp1
    cp2 = &cp2

    type CycleSlice []CycleSlice
    var cs = make(CycleSlice, 1)
    cs[0] = cs

    ch1, ch2 := make(chan int), make(chan int)
    var ch1ro <- chan int = ch1

    type mystring string

    var iface1, iface1Again, iface2 interface{} = &one, &oneAgain, &two

    for _, test := range []struct {
        x, y interface{}
        want bool
    }{
        // basic types
        {1, 1, true},
        {1, 2, false},
        {1, 1.0, false},
        {"foo", "foo", true},
        {"foo", "bar", false},
        {mystring("foo"), "foo", false},
        // slices
        {[]string{"foo"}, []string{"foo"}, true},
        {[]string{"foo"}, []string{"bar"}, false},
        {[]string{}, []string(nil), true},
        // slice cycles
        {cs, cs, true},
        // maps
        {
            map[string][]int{"foo": {1, 2, 3}},
            map[string][]int{"foo": {1, 2, 3}},
            true,
        },
        {
            map[string][]int{"foo": {1, 2, 3}},
            map[string][]int{"foo": {1, 2, 3, 4}},
            false,
        },
        {
            map[string][]int{},
            map[string][]int(nil),
            true,
        },
        // pointers
        {&one, &one, true},
        {&one, &two, false},
        {&one, &oneAgain, true},
        {new(bytes.Buffer), new(bytes.Buffer), true},
        // pointer cycles
        {cp1, cp1, true},
        {cp2, cp2, true},
        {cp1, cp2, true},
        // functions
        {(func())(nil), (func())(nil), true},
        {(func())(nil), func() {}, false},
        {func() {}, func() {}, false},
        // arrays
        {[...]int{1, 2, 3}, [...]int{1, 2, 3}, true},
        {[...]int{1, 2, 3}, [...]int{1, 2, 4}, false},
        // channels
        {ch1, ch1, true},
        {ch1, ch2, false},
        {ch1ro, ch1, false},
        // interfaces
        {&iface1, &iface1, true},
        {&iface1, &iface2, false},
        {&iface1Again, &iface1, true},
    } {
        if Equal(test.x, test.y) != test.want {
            t.Errorf("Equal(%v, %v) = %t", test.x, test.y, !test.want)
        }
    }
}

func ExampleEqual() {
    fmt.Println(Equal([]int{1, 2, 3}, []int{1, 2, 3}))
    fmt.Println(Equal([]string{"foo"}, []string{"bar"}))
    fmt.Println(Equal([]string(nil), []string{}))
    fmt.Println(Equal(map[string]int(nil), map[string]int{}))

    // cycle
    type link struct {
        value string
        tail *link
    }
    a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
    a.tail, b.tail, c.tail = b, a, c
    fmt.Println(Equal(a, a))
    fmt.Println(Equal(b, b))
    fmt.Println(Equal(c, c))
    fmt.Println(Equal(a, b))
    fmt.Println(Equal(a, c))

    // Output:
    // true
    // false
    // true
    // true
    // true
    // true
    // true
    // false
    // false
}
