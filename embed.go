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

package neng

import (
	"embed"
	"strings"
)

//go:embed embed/*
var efs embed.FS

// readEFS loads a word list from the embedded path.
// Returns an error if the file is not found.
func readEFS(path string) ([]string, error) {
	stream, err := efs.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(stream), "\n"), nil
}
