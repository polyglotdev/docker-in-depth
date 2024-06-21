# Managing Applications with Docker Compose

[Course Link](https://dominiquehallan-links.com/3VRsKy7)

## Docker Compose Overview

Docker Compose is a tool for defining and running multi-container Docker applications. With Compose, you use a YAML file to configure your application’s services. Then, with a single command, you create and start all the services from your configuration.

Compose works in all environments: production, staging, development, testing, as well as CI workflows. You can learn more about Compose in the [Compose file reference](https://docs.docker.com/compose/compose-file/).

Using Compose is basically a three-step process:

1. Define your app’s environment with a `Dockerfile` so it can be reproduced anywhere.
2. Define the services that make up your app in `docker-compose.yml` so they can be run together in an isolated environment.
3. Run `docker-compose up` and Compose starts and runs your entire app.
4. You can also run `docker-compose up -d` to run your services in the background.
5. You can also run `docker-compose down` to stop your services.
6. You can also run `docker-compose ps` to see the status of your services.
7. You can also run `docker-compose logs` to see the logs of your services.
8. You can also run `docker-compose exec` to run commands in your services.

## Docker Compose Installation

To install Docker Compose, you can follow the instructions in the [official documentation](https://docs.docker.com/compose/install/).

## Docker Compose Commands

Here are some useful Docker Compose commands:

- `docker-compose up` - Builds, (re)creates, starts, and attaches to containers for a service.
- `docker-compose up -d` - Builds, (re)creates, starts, and attaches to containers for a service in detached mode.
- `docker-compose down` - Stops and removes containers, networks, volumes, and images created by `up`.
- `docker-compose ps` - Lists containers.
- `docker-compose logs` - Views output from containers.
- `docker-compose exec` - Runs a command in a service.
- `docker-compose build` - Builds or rebuilds services.
- `docker-compose pull` - Pulls images for services.
- `docker-compose start` - Starts existing containers for a service.
- `docker-compose stop` - Stops existing containers without removing them.
- `docker-compose restart` - Restarts services.
- `docker-compose pause` - Pauses services.
- `docker-compose unpause` - Unpauses services.
- `docker-compose top` - Displays the running processes.
- `docker-compose events` - Gets real time events from containers.
- `docker-compose port` - Prints the public port for a port binding.
- `docker-compose config` - Validates and view the compose file.
- `docker-compose version` - Shows the Docker-Compose version information.
- `docker-compose help` - Shows help.

## Docker Compose File Reference

You can learn more about the Docker Compose file in the [official documentation](https://docs.docker.com/compose/compose-file/).

## Docker Compose Examples

Here are some examples of Docker Compose files:

- [awesome-compose](https://dominiquehallan-links.com/3xsyzIZ)
