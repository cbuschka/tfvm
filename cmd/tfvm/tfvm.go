package main

import (
	"github.com/cbuschka/tfvm/internal/cli"
	"github.com/cbuschka/tfvm/internal/log"
	"github.com/cbuschka/tfvm/internal/util"
	"os"
)

func main() {
	err := cli.Run(os.Args)
	if err != nil {
		util.Die(1, err.Error())
	}

	log.Infof("Exiting with code 0.")
	os.Exit(0)
}
