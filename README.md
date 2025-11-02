# Power Monitor

A simple application written in Go that send a notification if a client fails to check in with a server.

## Usage

### Compile the binary

`make build`

### Set up the configuration

Copy `example.yml` to `config.yml`, and update `config.yml` to contain appropriate values for sending email alerts.

### Run the server

`./power-monitor -server`

### Run the server in docker

```
docker-compose build
docker-compose up -d
```

### Run the client

`./power-monitor -client -hostname localhost -port 8000 -name test`

#### Systemd Service

You can run the power monitor client as a Systemd service by following these steps:

1. Compile the binary: `make build`
2. Create a script in the top-level directory called `run.sh`. Populate it with the command to run the client. For example:

```sh
#!/bin/bash
./power-monitor -client -hostname <server-hostname> -port <server-port> -name <client-name>
```

3. Make `run.sh` executable: `chmod +x run.sh`.
4. Run the `create-service.sh` script from the top-level directory: `./scripts/create-services.sh`. This will create, start, and enable the systemd service.

## Local Development

### Run tests

`make test`
