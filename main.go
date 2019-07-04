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

	// if runtime.GOOS == "linux" {
	// 	return ".so"
	// }

	return ".so"
}

func flags() string {
	// if runtime.GOOS == "darwin" {
	// 	return "-ldflags -s"
	// }

	if *fl {
		if runtime.GOOS == "windows" {
			return `-ldflags -s`
		}

		return `-ldflags "-s -w"`
	}

	return ""
}

func all() string {
	if *al {
		return "-a "
	}

	return ""
}

var (
	fname = flag.String("n", "", "file name")
	fl    = flag.Bool("f", false, "set -ld flags")
	al    = flag.Bool("a", false, "set build -a")

	dir = "./lib/"
)

func run(name string) {
	fmt.Println("build c-shared start...")

	str := "go build " + all() + flags() + " -o " + name + so() +
		" -buildmode=c-shared " + name + ".go"

	out, e, err := cmd.Run(str)
	// if err != nil {
	fmt.Println("cmd.Run err is: ", out, e, err)
	// }

	fmt.Println("end")
}

func main() {
	flag.Parse()

	run(*fname)
}
