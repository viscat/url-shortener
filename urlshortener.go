package main

import (
	"github.com/jessevdk/go-flags"
	"os"
	"urlshortener/cmd"
)

func main() {

	initLogger(true)
	parser := flags.NewNamedParser("urlshortener", flags.Default)
	parser.AddCommand("api", "Starts url shortener API", "", &cmd.ApiCommand{})

	_, err := parser.Parse()
	if err != nil {
		if flagErr, ok := err.(*flags.Error); ok && flagErr.Type != flags.ErrHelp {
			parser.WriteHelp(os.Stdout)
		}

		os.Exit(1)
	}
}
