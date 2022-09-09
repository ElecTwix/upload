package main

import (
	"errors"
	"log"
	"os"

	"github.com/pterm/pterm"
)

func checklaunch() (Args []string, err error) {
	args := os.Args
	if len(args) == 3 {
		return args, nil
	}

	return nil, errors.New("exected 2 commman line arguments")
}

func checkErr(err error, section string) {
	if err != nil {

		log.Fatalf(pterm.NewStyle(pterm.BgWhite, pterm.BgRed, pterm.Bold).Sprintf("%s: %s", section, err.Error()))
	}
}
