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

// Package common contains functions shared by scripts.
package common

import (
	"crypto/sha256"
	"fmt"
	"os"
	"slices"
	"strings"
)

// ReadFile reads a file at OS path, splits its content into lines
// and returns the resulting slice.
func ReadFile(path string) ([]string, error) {
	stream, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(stream), "\n")

	if len(lines[len(lines)-1]) == 0 {
		return lines[:len(lines)-1], nil
	}
	return lines, nil
}

// WriteFile writes lines to the specified path, optionally sorting them.
// Returns SHA256 checksum of the data written and errors related
// to file handling.
func WriteFile(path string, lines []string, sort bool) (string, error) {
	if sort {
		slices.Sort(lines)
	}

	data := []byte(strings.Join(lines, "\n"))

	if len(data) > 0 && data[len(data)-1] != '\n' {
		data = append(data, '\n')
	}

	csum := fmt.Sprintf("%x  %s", sha256.Sum256(data), path)

	return csum, os.WriteFile(path, data, 0644)
}
