// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

package daemon

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var flagBgJob = flag.Bool("b", false, "run as background job")

func usage() {
	fmt.Fprintf(os.Stderr, "usage: fruitd [flags] <fruits>\n")
	flag.PrintDefaults()
}

func run(fruits ...string) {
	msg := "Starting fruit daemon"
	if *flagBgJob {
		msg += " as background job"
	}
	log.Printf("%s for %s...\n", msg, strings.Join(fruits, ", "))
}

func Main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "error: at least one fruit is required")
		os.Exit(2)
	}

	run(flag.Args()...)
}
