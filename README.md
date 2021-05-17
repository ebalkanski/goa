# Service Overview

Just a simple app go play with GOA framework and MongoDB
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

Create `.env` and adjust it if needed.
```bash
cp .env.dist .env
```

Start the containers.
```bash
docker-compose up -d 
```

The command will start one Docker container - the `goa` service container listening 
for requests on port 8080 and accessible via browser on `http://localhost:8080` 
and one MongoDB container listening on port 27017.

There is a code watcher running in the container, so if you change something in the code, the service will be restarted automatically and you will be able to see the changes immediately in the browser.

You can see the logs of the running container by executing:
```bash
docker logs -f goa
```

Test it by opening [http://localhost:8080](http://localhost:8080) in a browser or use Curl from command line:
```bash
curl -i -H "Content-Type:application/json" localhost:8080
```

The Swagger UI for the API is available at [http://localhost:8080/swagger-ui](http://localhost:8080/swagger-ui)