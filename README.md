# go-iterable

[![GoDoc](https://godoc.org/github.com/r-evolve/go-iterable?status.svg)](https://godoc.org/github.com/r-evolve/go-iterable)

## Summary
go-iterable is a Go library for working with iterables.
It is inspired by the Java 8 Stream API.

## Installation
```bash
go get github.com/r-evolve/go-iterable
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/r-evolve/go-iterable"
)

func main() {
    // Create an iterable from a slice
    stream := go_iterable.New([]int{1, 2, 3, 4, 5})

    // Filter the iterable
    stream = stream.Filter(func(i int) bool {
        return i > 2
    })

    // Map the iterable
    stream = stream.Map(func(i int) int {
        return i * 2
    })

    // Reduce the iterable
    sum := stream.Reduce(func(a, b int) int {
        return a + b
    }, 0)

    // Print the result
    fmt.Println(sum)
}
```

## Usage alternative

```go
package main

import (
    "fmt"
    "github.com/r-evolve/go-iterable"
)

func main() { 
    // Create an iterable from a slice
    stream := go_iterable.New([]int{1, 2, 3, 4, 5})

    // Print the result
    fmt.Println(stream.Filter(func(i int) bool {
        return i > 2
    }).Map(func(i int) int {
        return i * 2
    }).Reduce(func(a, b int) int {
        return a + b
    }, 0))
}
```

## Other examples

```go
package main

import (
    "fmt"
    "github.com/r-evolve/go-iterable"
)

func main() {
    // Create channels from single values
    _ = go_iterable.Of(1, 2, 3, 4, 5).ToChannel()
}

```

## Documentation
[![GoDoc](https://godoc.org/github.com/r-evolve/go-iterable?status.svg)](https://godoc.org/github.com/r-evolve/go-iterable)