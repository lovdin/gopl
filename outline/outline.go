package main

import (
    "fmt"
    "os"
    "net/http"
    "strings"

    "golang.org/x/net/html"
)

func main() {
    for _, url := range os.Args[1:] {
        outline(url)
    }
}

func outline(url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("getting %s: %s", url, resp.Status)
    }

    ct := resp.Header.Get("Content-Type")
    if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
        return fmt.Errorf("%s has type %s, not text/html", url, ct)
    }

    doc, err := html.Parse(resp.Body)
    if err != nil {
        return err
    }

    forEachNode(doc, startElement, endElement)

    return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
    if pre != nil {
        pre(n)
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, pre, post)
    }

    if post != nil {
        post(n)
    }
}

var depth int

func startElement(n *html.Node) {
    if n.Type == html.ElementNode {
        var attr string
        for _, a := range n.Attr {
            attr += " " + a.Key
        }

        fmt.Printf("%*s<%s%s>\n", depth * 2, "", n.Data, attr)
        depth++
    }
}

func endElement(n *html.Node) {
    if n.Type == html.ElementNode {
        depth--
        fmt.Printf("%*s</%s>\n", depth * 2, "", n.Data)
    }
}
