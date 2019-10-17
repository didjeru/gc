package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var pattern string
	flag.StringVar(&pattern, "e", "", "RegExp pattern")
	flag.Parse()

	if pattern == "" && len(flag.Args()) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s [-e] pattern\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		var matched bool
		var err error

		if pattern != "" {
			matched, err = regexp.MatchString(pattern, line)
		} else {
			searchValue := os.Args[len(os.Args)-1]
			matched = strings.Contains(line, searchValue)
		}

		if err != nil {
			panic(err)
		}

		if matched {
			fmt.Println(line)
		}
	}
}
