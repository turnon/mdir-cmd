package main

import (
	"errors"
	"os"
	"strconv"

	"github.com/turnon/mdir"
	"github.com/urfave/cli/v2"
)

const helpTmpl = `NAME:
  {{.Name}} - {{.Usage}}

USAGE:
  mdir-cmd /src /dest 2 3 4
{{if .VisibleFlags}}
GLOBAL OPTIONS:
{{range .VisibleFlags}}  {{.}}
{{end}}
{{end}}
`

func main() {
	cli.AppHelpTemplate = helpTmpl

	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:  "f",
			Usage: "force to mv/cp files",
		},
		&cli.BoolFlag{
			Name:  "c",
			Usage: "copy instead of move",
		},
		&cli.BoolFlag{
			Name:  "p",
			Usage: "show progress bar",
		},
	}

	app := &cli.App{
		Name:                   "mdir-cmd",
		Usage:                  "cmd of mdir",
		Version:                "v1.0.0",
		Action:                 action,
		Flags:                  flags,
		UseShortOptionHandling: true,
	}

	err := app.Run(os.Args)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
	}
}

func action(c *cli.Context) error {
	if c.NArg() < 3 {
		return errors.New(c.App.Usage)
	}

	args := c.Args().Slice()

	segments := make([]int, 0, len(args)-2)
	for _, seg := range args[2:] {
		i, err := strconv.Atoi(seg)
		if err != nil {
			return err
		}
		segments = append(segments, i)
	}

	cmd := mdir.Cmd{
		Src:      args[0],
		Dest:     args[1],
		Segments: segments,
		Progress: c.Bool("p"),
		Force:    c.Bool("f"),
		CopyFile: c.Bool("c"),
	}

	return cmd.MvFiles()
}
