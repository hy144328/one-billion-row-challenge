package main

import (
	"bytes"
)

type Register[T any] struct {
	Key []byte
	Value T
}

type BytesMap[T any] struct {
	noRegisters uint32
	registers []Register[T]
}

func NewBytesMap[T any](noRegisters int) *BytesMap[T] {
	if noRegisters & (noRegisters - 1) != 0 {
		panic("not power of 2")
	}

	return &BytesMap[T]{
		noRegisters: uint32(noRegisters),
		registers: make([]Register[T], noRegisters),
	}
}

func (m *BytesMap[T]) GetOrCreate(k []byte) (*T, bool) {
	h := calculateHash(k)

	for i := h; i < h + m.noRegisters; i++ {
		idx := i & (m.noRegisters - 1)

		if m.registers[idx].Key == nil {
			m.registers[idx].Key = make([]byte, len(k))
			copy(m.registers[idx].Key, k)
			return &m.registers[idx].Value, false
		}

		if bytes.Equal(m.registers[idx].Key, k) {
			return &m.registers[idx].Value, true
		}
	}

	return nil, false
}

func (m *BytesMap[T]) ToMap() map[string]*T {
	res := make(map[string]*T, maxCities)

	for i := range m.registers {
		if m.registers[i].Key != nil {
			res[string(m.registers[i].Key)] = &m.registers[i].Value
		}
	}

	return res
}

func calculateHash(bs []byte) uint32 {
	var res uint32 = 2166136261

	for _, b := range bs {
		res ^= uint32(b)
		res *= 16777619
	}

	return res
}
