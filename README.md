# docker-volume
Command line application to manage volumes of docker containers

# Install

```
go get -u github.com/dictybase-docker/docker-volume
```

# Usage

## Commands

```
    NAME:
       docker-volume - Command line app to manage volumes of docker containers

    USAGE:
       docker-volume [global options] command [command options] [arguments...]

    VERSION:
       1.0.0

    COMMANDS:
       export	Export a gzipped tar stream
       help, h	Shows a list of commands or help for one command
       
    GLOBAL OPTIONS:
       --host, -H 'unix:///var/run/docker.sock'	Endpoint for docker host [$DOCKER_HOST]
       --help, -h					show help
       --version, -v				print the version
       
```

## Subcommands
### Export

```
NAME:
   export - Export a gzipped tar stream

USAGE:
   command export [command options] [arguments...]

DESCRIPTION:
   

OPTIONS:
   --pause, -p		Pause running container before export
   --output, -o 	Where the output will be written, default is stdout
   
```

## Examples
Export the content of a data volume folder as a tar.gz archive

```
docker-volume export -o bohr.tar.gz modest_bohr:/data
```

The output can also be redirected

```
docker-volume export pensive_euclid:/config > pensive.tar.gz
```

