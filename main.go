package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/jlecour/rrdbeat/beater"
)

func main() {
	err := beat.Run("rrdbeat", "", beater.New())
	if err != nil {
		os.Exit(1)
	}
}
