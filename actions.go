package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/fsouza/go-dockerclient"
	"gopkg.in/codegangsta/cli.v1"
)

func Validate(c *cli.Context) error {
	if len(c.Args()) == 0 {
		return fmt.Errorf("container name is missing")
	}
	return nil
}

func ExportAction(c *cli.Context) {
	if err := Validate(c); err != nil {
		log.Fatal(err)
	}
	client, err := GetClient(c.GlobalString("host"))
	if err != nil {
		log.Fatal(err)
	}
	// The argument is given in container:volume format
	cfmt := strings.Split(c.Args()[0], ":")
	// Get the container id from name
	id, err := GetContainerId(client, cfmt[0])
	if err != nil {
		log.Fatal(err)
	}
	// check if the volume exists
	cinfo, err := client.InspectContainer(id)
	if err != nil {
		log.Fatal(err)
	}
	volfound := false
	for cvol, _ := range cinfo.Volumes {
		if cvol == cfmt[1] {
			volfound = true
		}
	}
	if !volfound {
		log.Fatalf("Given volume %s is not found in container %s\n", cfmt[1], cfmt[0])
	}
	// copy content of volume from container
	var w io.Writer
	if c.IsSet("output") {
		f, err := os.Create(c.String("output"))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		w = f
	} else {
		w = os.Stdout
	}
	gw := gzip.NewWriter(w)
	defer gw.Close()
	err = client.CopyFromContainer(docker.CopyFromContainerOptions{OutputStream: gw, Container: id, Resource: cfmt[1]})
	if err != nil {
		log.Fatal(err)
	}
}
