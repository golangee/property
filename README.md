# property [![GoDoc](https://godoc.org/github.com/golangee/property?status.svg)](http://godoc.org/github.com/golangee/property)
Package property provides some basic properties which can be observed.
Providing type safe properties in Go is impossible without template or generic types. Therefore, we
provide a "generic" type-unsafe base implementation and a few type safe wrappers for primitives.

## Roadmap
It is planned to replace this with a single implementation using type parameters (aka Generics) as soon
as they land in Go.