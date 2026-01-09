package main

type Statistics[T any] struct {
	cnt int
	max T
	min T
	sum T
}
