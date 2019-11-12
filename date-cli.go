package main

import (
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "Date CLI"
	app.Usage = "date-cli [timestamp] [output]"

	app.Action = func(c *cli.Context) error {

		loc, err := time.LoadLocation("GMT")
		if err != nil {
			panic(err.Error())
		}
		time.Local = loc

		t, err := dateparse.ParseLocal(c.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}

		outFmt := c.Args().Get(1)
		if outFmt == "u" || outFmt == "unix" {
			fmt.Println(t.Unix())
		} else {
			fmt.Println(t)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
