package main

import (
	"bytes"
)

type Register[T any] struct {
	KeyLen int
	Key [100]byte
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

		if klen := m.registers[idx].KeyLen; klen == 0 {
			m.registers[idx].KeyLen = len(k)
			copy(m.registers[idx].Key[:], k)
			return &m.registers[idx].Value, false
		} else if bytes.Equal(m.registers[idx].Key[:klen], k) {
			return &m.registers[idx].Value, true
		}
	}

	panic("registers full")
}

func (m *BytesMap[T]) ToMap() map[string]*T {
	res := make(map[string]*T, maxCities)

	for i := range m.registers {
		if klen := m.registers[i].KeyLen; klen > 0 {
			res[string(m.registers[i].Key[:klen])] = &m.registers[i].Value
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
