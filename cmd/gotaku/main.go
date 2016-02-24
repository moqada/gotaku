package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/moqada/gotaku"
)

var (
	listing bool
)

func init() {
	flag.BoolVar(&listing, "l", false, "Listing GYOTAKU urls")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <URL>\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
		fmt.Println("")
		fmt.Fprintln(os.Stderr, "Examples:")
		fmt.Fprintln(os.Stderr, "  gotaku http://example.com/     Take GYOTAKU of specific url")
		fmt.Fprintln(os.Stderr, "  gotaku -l http://example.com/  List GYOTAKU urls for specific url")
	}
	flag.Parse()
}

func list(target string) error {
	urls, err := gotaku.List(target)
	if err != nil {
		return err
	}
	for _, u := range urls {
		fmt.Println(u)
	}
	return nil
}

func take(target string) error {
	url, err := gotaku.Take(target)
	if err != nil {
		return err
	}
	fmt.Println(url)
	return nil
}

func run() error {
	args := flag.Args()
	if len(args) < 1 {
		return fmt.Errorf("Target URL does not set.")
	}
	if listing {
		return list(args[0])
	}
	return take(args[0])
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
