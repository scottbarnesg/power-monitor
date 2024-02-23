# Power Monitor

A simple application written in Go that send a notification if a client fails to check in with a server.

## Usage

### Compile the binary

`make build`

### Set up the configuration

Copy `example.yml` to `config.yml`, and update `config.yml` to contain appropriate values for sending email alerts.

### Run the server

`./power-monitor -server`

### Run the client

`./power-monitor -client -hostname localhost -port 8000 -name test`

## Local Development

### Run tests

`make test`
