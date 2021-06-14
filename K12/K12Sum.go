// (K12Sum.go) - Contains the K12 Hash functions of ./mfc
// Copyright (C) 2021 MaxflowO2, the only author of Max Flow Chain
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//
// Huge shoutout to David Wong @ https://www.cryptologie.net/contact
// We won't touch the orginals, but here's all the "SumXXX" values
// for you brother! Thank you for showing "us" aka MaxflowO2 the way!

package K12

// nil custom string, change if you wish
var customString []byte

// Sum224([]byte)
// Returns a []byte, len 28, K12 Hash
func Sum224(input []byte) []byte {
        value := make([]byte, 28)
        K12Sum(customString, input, value)
        return value
}

// Sum256([]byte)
// Returns a []byte, len 32, K12 Hash
func Sum256(input []byte) []byte {
	value := make([]byte, 32)
	K12Sum(customString, input, value)
	return value
}

// Sum384([]byte)
// Returns a []byte, len 48, K12 Hash
func Sum384(input []byte) []byte {
        value := make([]byte, 48)
        K12Sum(customString, input, value)
        return value
}

// Sum512([]byte)
// Return a []byte, len 64, K12 Hash
func Sum512(input []byte) []byte {
	value := make([]byte, 64)
	K12Sum(customString, input, value)
	return value
}
