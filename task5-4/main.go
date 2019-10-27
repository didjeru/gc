package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func copyFile(source, destination string) (err error) {
	src, err := os.Open(source)
	if err != nil {
		return err
	}

	fi, err := src.Stat()
	if err != nil {
		return err
	}
	flag := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	perm := fi.Mode() & os.ModePerm
	dst, err := os.OpenFile(destination, flag, perm)
	if err != nil {
		return err
	}

	_, err = io.Copy(dst, src)
	if err != nil {
		os.Remove(destination)
		return err
	}

	defer func() {
		err := src.Close()
		printError(err)
		err = dst.Close()
		printError(err)
	}()

	return nil
}

func main() {
	var src, dst string
	flag.StringVar(&src, "src", "", "source file")
	flag.StringVar(&dst, "dst", "", "destination file")
	flag.Parse()
	if src == "" || dst == "" {
		flag.Usage()
		os.Exit(1)
	}

	err := copyFile(src, dst)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("copied %q to %q\n", src, dst)
}
