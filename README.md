# Service Overview

Just a simple app go play with GOA framework
- `http://localhost:8080`

## Prerequisites

* Install `git`
* [Install Go](https://golang.org/doc/install) and set the
[`$GOPATH` environment variable](https://github.com/golang/go/wiki/SettingGOPATH)
* Install `docker` and `docker-compose`
* Clone the `goa` repo

The repo must be cloned in the `$GOPATH/src/github.com/ebalkanski/goa/` directory. 
You must create it if it doesn't exist.

```bash
mkdir -p $GOPATH/src/github.com/ebalkanski/goa/
cd $GOPATH/src/github.com/ebalkanski/goa/
git clone https://github.com/ebalkanski/goa.git .
```

## Run with Docker

Build the Docker image.
```bash
docker build -t goa -f Dockerfile.dev .
```

Start the `goa` container.
```bash
docker run -p 8080:8080 -v $GOPATH/src/github.com/ebalkanski/goa/:/go/src/github.com/ebalkanski/goa/ --name goa -d goa
```

The command will start one Docker container - the `goa` service container listening 
for requests on port 8080 and accessible via browser on `http://localhost:8080`.

There is a code watcher running in the container, so if you change something in the code, the service will be restarted automatically and you will be able to see the changes immediately in the browser.

You can see the logs of the running container by executing:
```bash
docker logs -f goa
```

Test it by opening [http://localhost:8080](http://localhost:8080) in a browser or use Curl from command line:
```bash
curl -i -H "Content-Type:application/json" localhost:8080
```