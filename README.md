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
	iterable := go_iterable.NewIterable([]int{1, 2, 3, 4, 5})

	// Filter the iterable
	iterable = iterable.Filter(func(i int) bool {
		return i > 2
	})

	// Map the iterable
	iterable = iterable.Map(func(i int) int {
		return i * 2
	})

	// Reduce the iterable
	sum := iterable.Reduce(func(a, b int) int {
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
    iterable := go_iterable.NewIterable([]int{1, 2, 3, 4, 5})

    // Print the result
    fmt.Println(iterable.Filter(func(i int) bool {
	return i > 2
    }).Map(func(i int) int {
        return i * 2
    }).Reduce(func(a, b int) int {
        return a + b
    }, 0))
}
```

## Documentation
[![GoDoc](https://godoc.org/github.com/r-evolve/go-iterable?status.svg)](https://godoc.org/github.com/r-evolve/go-iterable)