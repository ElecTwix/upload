package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ElecTwix/uploader"
	"github.com/pterm/pterm"
)

func main() {
	args, err := checklaunch()
	checkErr(err, "launch")

	file, err := os.OpenFile(args[1], os.O_RDWR, 0644)

	checkErr(err, "open file")

	var site uploader.Site
	switch args[2] {

	case "anon":
		site = uploader.AnonFiles
		break
	case "bay":
		site = uploader.BayFiles
		break
	default:
		log.Fatalf(pterm.NewStyle(pterm.FgWhite, pterm.BgRed, pterm.Bold).Sprintf("'%s' is not valid site", args[2]))
	}

	uploadspinner, _ := pterm.DefaultSpinner.Start("Uploading ...")

	respond, err := uploader.Upload(site, file)

	if err != nil {
		uploadspinner.Info(err.Error())
	}

	uploadspinner.Success(fmt.Sprintf("Short: %s  /  Long: %s", respond.Data.File.Url.Short, respond.Data.File.Url.Full))

}
