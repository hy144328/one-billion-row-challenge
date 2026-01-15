package main

import (
	"bytes"
	"hash/maphash"
)

type Register[T any] struct {
	KeyLen int
	Hash   uint64
	Key    [100]byte
	Value  T
}

type BytesMap[T any] struct {
	noRegisters uint64
	registers   []Register[T]
	seed        maphash.Seed
}

func NewBytesMap[T any](noRegisters int) *BytesMap[T] {
	if noRegisters&(noRegisters-1) != 0 {
		panic("not power of 2")
	}

	return &BytesMap[T]{
		noRegisters: uint64(noRegisters),
		registers:   make([]Register[T], noRegisters),
		seed:        maphash.MakeSeed(),
	}
}

func (m *BytesMap[T]) GetOrCreate(k []byte) (*T, bool) {
	h := maphash.Bytes(m.seed, k)

	for i := h; i < h+m.noRegisters; i++ {
		idx := i & (m.noRegisters - 1)

		if klen := m.registers[idx].KeyLen; klen == 0 {
			m.registers[idx].KeyLen = len(k)
			m.registers[idx].Hash = h
			copy(m.registers[idx].Key[:], k)
			return &m.registers[idx].Value, false
		} else if h == m.registers[idx].Hash && bytes.Equal(m.registers[idx].Key[:klen], k) {
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
