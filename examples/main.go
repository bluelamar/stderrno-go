package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bluelamar/stderrno-go/stderrno"
)

// openFile will open the specified file for reading.
func openFile(fname string) error {
	f, err := os.OpenFile(fname, os.O_RDONLY, 0)
	if err != nil {
		log.Printf("openFile returns internal error=%v\n", err)

		return fmt.Errorf("%v: %w", err, stderrno.ENOENT)
	}

	f.Close()
	return nil
}

// httpGetOpen performs a GET against the specified url.
func httpGetOpen(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("http.Get returns internal error=%v\n", err)

		return fmt.Errorf("%v: %w", err, stderrno.EHOSTUNREACH)
	}

	_ = resp.Body.Close()

	return nil
}

func main() {
	log.Printf("Open non-existent file:")
	err := openFile("judy-blue-eyes")
	log.Printf("openFile returned error=%v", err)
	if err != nil {
		unwerr := errors.Unwrap(err)
		log.Printf("unwrapped error=%v", unwerr)
	}

	if errors.Is(err, stderrno.ENOENT) {
		log.Printf("SUCCESS: openFile matched the expected std error: %v", err)
	} else {
		log.Printf("openFile returned unexpected error: %v", err)
	}

	log.Printf("-----")

	log.Printf("Open non-existent http site:")
	err = httpGetOpen("http://brown-eyed-girl")
	if err != nil {
		unwerr := errors.Unwrap(err)
		log.Printf("unwrapped error=%v", unwerr)
	}

	if errors.Is(err, stderrno.EHOSTUNREACH) {
		log.Printf("SUCCESS: httpGetOpen matched the expected std error: %v", err)
	} else {
		log.Printf("httpGetOpen returned unexpected error: %v", err)
	}

	log.Printf("-----")
}
