package main

import (
	"fmt"

	"github.com/project-machine/mos/pkg/mosconfig"
	"github.com/urfave/cli"
)

var installCmd = cli.Command{
	Name:   "install",
	Usage:  "install a new mos system",
	Action: doInstall,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "f, file",
			Usage: "File from which to read the install manifest",
			Value: "./install.yaml",
		},
		cli.StringFlag{
			Name:  "config-dir, c",
			Usage: "Directory where mos config is found",
			Value: "/config",
		},
		cli.StringFlag{
			Name:  "atomfs-store, a",
			Usage: "Directory under which atomfs store is kept",
			Value: "/atomfs-store",
		},
	},
}

func doInstall(ctx *cli.Context) error {
	// Expect config, scratch-writes, and atomfs-store to exist
	store := ctx.String("atomfs-store")
	if !mosconfig.PathExists(store) {
		return fmt.Errorf("atomfs store not found")
	}

	config := ctx.String("config-dir")
	if !mosconfig.PathExists(config) {
		return fmt.Errorf("mos config directory not found")
	}

	if err := mosconfig.InitializeMos(store, config, ctx.String("file")); err != nil {
		return err
	}

	return nil
}
