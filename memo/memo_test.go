package memo_test

import (
    "testing"

    "github.com/lovdin/gopl/memo"
    "github.com/lovdin/gopl/memotest"
)

var httpGetBody = memotest.HTTPGetBody

// $ go test -v github.com/lovdin/gopl/memo
func Test(t *testing.T) {
    m := memo.New(httpGetBody)
    memotest.Sequential(t, m)
}

// $ go test -run=TestConcurrent -race -v github.com/lovdin/gopl/memo
func TestConcurrent(t *testing.T) {
    m := memo.New(httpGetBody)
    memotest.Concurrent(t, m)
}
