package main

import (
	"fmt"
	"github.com/keitarou/go-jj/command"
	"github.com/mitchellh/cli"
	"os"
)

func main() {

	c := cli.NewCLI("jj", "0.0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"to_timestamp": func() (cli.Command, error) {
			return &command.ToTimestampCommand{}, nil
		},
		"to_date": func() (cli.Command, error) {
			return &command.ToDateCommand{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(exitStatus)
}
