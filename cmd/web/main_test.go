package main

import (
	"log"
	"testing"
)

func TestRun(t *testing.T) {
	_, err := run()
	if err != nil {
		log.Fatal("Test did not pass: please check 'run' function in package main")
	}
}
