// Copyright 2016 The go-vgo Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/vcaesar/gob/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/go-vgo/gt/cmd"
)

func so() string {
	if runtime.GOOS == "darwin" {
		return ".dylib"
	}

	if runtime.GOOS == "windows" {
		return ".dll"
	}

	if runtime.GOOS == "linux" {
		return ".so"
	}

	return ""
}

func flags() string {
	if runtime.GOOS == "darwin" {
		return "-ldflags -s"
	}

	return ""
}

var (
	fname = flag.String("n", "", "file name")

	dir = "./lib/"
)

func run(name string) {
	fmt.Println("build c-shared start...")

	str := "go build " + flags() + " -o " + name + so() +
		" -buildmode=c-shared " + name + ".go"

	cmd.Run(str)
	fmt.Println("end")
}

func main() {
	flag.Parse()

	run(*fname)
}
