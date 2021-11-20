// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package promexp

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	flagScrapeEndpoint = flag.String("s", "", "the endpoint for scraping")
	flagPromURL        = flag.String("p", "", "the Prometheus connection URL")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: fruitpromexp [flags]\n")
	flag.PrintDefaults()
}

func export() {
	log.Printf("Exporting data from %q to %q...\n", *flagScrapeEndpoint, *flagPromURL)
}

func Main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NFlag() != 2 {
		fmt.Fprintln(os.Stderr, "error: scrape endpoint and Prometheus connection URL are required")
		os.Exit(1)
	}
	if *flagPromURL == "" || *flagScrapeEndpoint == "" {
		fmt.Fprintln(os.Stderr, "error: scrape endpoint or Prometheus connection URL must not be empty")
		os.Exit(1)
	}

	export()
}
