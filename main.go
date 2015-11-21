package main

import (
	"log"
	"os"

	"github.com/humboldtux/cisgo/diff"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	diff.Run()
}
