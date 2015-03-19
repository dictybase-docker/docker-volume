package main

import (
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
)

func GetClient(host string) (*docker.Client, error) {
	c, err := docker.NewClient(host)
	if err != nil {
		return c, fmt.Errorf("Error connecting to docker daemon: %s\n", err)
	}
	return c, nil
}

func GetContainerId(client *docker.Client, name string) (string, error) {
	allconts, err := client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		return "", err
	}
	// add forward slash to name as that's the way api reports
	name = fmt.Sprintf("/%s", name)
	for _, ct := range allconts {
		for _, n := range ct.Names {
			if n == name {
				return ct.ID, nil
			}
		}
	}
	return "", fmt.Errorf("Could not find id for container %s\n", name)
}
