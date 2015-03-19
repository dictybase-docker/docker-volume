package main

import (
	"os"

	"gopkg.in/codegangsta/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "docker-volume"
	app.Version = "1.0.0"
	app.Usage = "Command line app to manage volumes of docker containers"
	app.Author = "Siddhartha Basu"
	app.Email = "siddhartha-basu@northwestern.edu"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "host, H",
			Usage:  "Endpoint for docker host",
			Value:  "unix:///var/run/docker.sock",
			EnvVar: "DOCKER_HOST",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "export",
			Usage:  "Export a gzipped tar stream",
			Action: ExportAction,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "pause, p",
					Usage: "Pause running container before export",
				},
				cli.StringFlag{
					Name:  "output, o",
					Usage: "Where the output will be written, default is stdout",
				},
			},
		},
	}
	app.Run(os.Args)
}
