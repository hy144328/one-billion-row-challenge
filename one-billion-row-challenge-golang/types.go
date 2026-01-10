package main

type Statistics[T any] struct {
	Cnt int
	Max T
	Min T
	Sum T
}
