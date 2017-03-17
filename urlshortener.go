package main

import (
	"github.com/jessevdk/go-flags"
	"urlshortener/cmd"
	"os"
)

func main() {

	parser := flags.NewNamedParser("urlshortener", flags.Default)
	parser.AddCommand("api", "Starts url shortener API", "", &cmd.ApiCommand{})
	parser.AddCommand("start", "Starts url shortener service", "", &cmd.StartCommand{})

	_, err := parser.Parse()
	if err != nil {
		if flagErr, ok := err.(*flags.Error); ok && flagErr.Type != flags.ErrHelp {
			parser.WriteHelp(os.Stdout)
		}

		os.Exit(1)
	}
}
