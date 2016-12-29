package memo_test

import (
    "testing"

    "github.com/lovdin/gopl/memo2"
    "github.com/lovdin/gopl/memotest"
)

var httpGetBody = memotest.HTTPGetBody

// $ go test -v github.com/lovdin/gopl/memo
func Test(t *testing.T) {
    m := memo.New(httpGetBody)
    defer m.Close()
    memotest.Sequential(t, m)
}

// $ go test -run=TestConcurrent -race -v github.com/lovdin/gopl/memo
func TestConcurrent(t *testing.T) {
    m := memo.New(httpGetBody)
    defer m.Close()
    memotest.Concurrent(t, m)
}
