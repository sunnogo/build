// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The go1.8.4 command runs the go command from go1.8.4.
//
// To install, run:
//
//     $ go get github.com/sunnogo/build/version/go1.8.4
//     $ go1.8.4 download
//
// And then use the go1.8.4 command as if it were your normal go
// command.
//
// See the release notes at https://golang.org/doc/go1.8
//
// File bugs at https://golang.org/issues/new
package main

import "github.com/sunnogo/build/version"

func main() {
	version.Run("go1.8.4")
}
