package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var flagVerbose = flag.Bool("v", false, "enable verbose output")

func usage() {
	fmt.Fprintf(os.Stderr, "usage: fruitctl [flags] <fruits>\n")
	flag.PrintDefaults()
}

func tasty(fruits ...string) {
	msg := fmt.Sprintf("Washing %d tasty fruits", len(fruits))
	if *flagVerbose {
		msg = fmt.Sprintf("Tasty fruits: %s!", strings.Join(fruits, ", "))
	}
	fmt.Println(msg)
}

func Main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "error: at least one fruit is required")
		os.Exit(2)
	}

	tasty(flag.Args()...)
}
