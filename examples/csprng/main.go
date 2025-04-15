// neng -- Non-Extravagant Name Generator
// Copyright (C) 2024  Wojciech Głąb (github.com/Zedran)
//
// This file is part of neng.
//
// neng is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3 only.
//
// neng is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with neng.  If not, see <https://www.gnu.org/licenses/>.

// Example of using crypto/rand as random number generator for neng.
package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/Zedran/neng"
)

// source implements rand.Source.
type source struct {
	// Internal buffer holding random bytes read from random device
	buffer []byte
}

// Uint64 is defined by rand.Source. It reads bytes from random device
// and combines them into a random uint64.
func (s source) Uint64() uint64 {
	crand.Read(s.buffer)
	return binary.LittleEndian.Uint64(s.buffer)
}

// newSource constructs byte slice for the new source object.
func newSource() source {
	const U64_BYTE_COUNT int = 8
	return source{make([]byte, U64_BYTE_COUNT)}
}

func main() {
	src := rand.New(newSource())

	gen, err := neng.DefaultGenerator(src)
	if err != nil {
		log.Fatal(err)
	}

	for range 5 {
		phrase, err := gen.Phrase("%a %n that %Nv")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(phrase)
	}
}
