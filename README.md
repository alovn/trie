# trie

Trie tree by golang

## example

```go
package main

import (
    "fmt"
    "github.com/alovn/trie"
)

func main() {
    //creat a trie
    t := trie.New()

    //add keys
    t.Add("hello")

    //find keys
    exists := t.Find("hello")
    fmt.Println(exists) //true

    //remove keys
    t.Remove("hello")

    exists = t.Find("hello")
    fmt.Println(exists) //false
}
```
