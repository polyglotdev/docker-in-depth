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

## How to create a Docker Compose file Using YAML

## Agenda

- YAML
- Root elements
  - Version
  - Services
  - Networks
  - Volumes
- Special Topics

## YAML

- YAML is a human-readable data serialization standard that can be used in conjunction with all programming languages and is often used to write configuration files.
- uses `yaml` or `yml` file extension.
- Extensively used in configuration files.

## YAML Data Types

- Integers
  - 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
- Strings
  - `This is a string`
- Nulls
  - `~`
- Booleans
  - `true` or `false`
- Collections
  - Mappings
    - `key: value`
  - Nested Mappings
  - `key: value`
    - `key: value`
    - `key: value`
  - Inline Syntax
    - `key: { key: value, key: value }`
  - Sequences
    - `- item1`
    - `- item2`
    - `- item3`
    - `- item4`

## Compose Files

## Root Elements

- `version`: The version of the Compose file format being used.
- `services`: services in your application
- `networks`: networks your app uses
- `volumes`: volumes your app uses
- `secrets`: secrets your app uses
- `configs`: configs your app uses
- `extensions`: extensions your app uses

## Compose File Service Configuration

| Configuration             | docker run                  | Compose file key |
| ------------------------- | --------------------------- | ---------------- |
| **Specify image**         | `required arg`              | ` image`         |
| **Specify volume**        | `-v`, `--volume`, `--mount` | `volumes`        |
| **Publish ports on host** | `-p`                        | `ports`          |
| **Environment variables** | `-e`,`--env`                | `environment`    |
| **Logging**               | `--log-driver`, `--log-opt` | `logging`        |
| **Security options**      | `--security-opt`            | `security_opt`   |

## Compose file service Caveats

- some config only apply to swarm mode
- some are specified on the command line

## Compose File Dependencies

- `depends_on`: Express dependency between services
- `links`: Express dependency between services

## Compose File Service Examples

```bash
docker rn --name app-cache -d redis
```

```yaml
version: "3.8"
services:
  app-cache:
    image: redis
```

## Compose File Volumes

- `volumes`: named volumes, anonymous volumes, bind mounts
- Allows naming and sharing of volumes
- Supports external volumes

## Compose File Networks

- `networks`: bridge, overlay, macvlan, none
- One network default by default

## Special Topics

- use shell environment variables in your Compose file
- can use `$variable` or `${variable}` syntax

## Docker Compose CLI

- Features
- Installation
- Usage

## Docker Compose CLI Features

- run multiple isolated environments on a single host
- parallel execution model
- Compose file change detection
- Open Source

## Compose Usage

- `docker compose [options] [COMMAND] [ARGS...]`
- `docker compose --help`
- Looks for `docker-compose.yml` or `docker-compose.yaml` file by default in the current directory
- Projects to represent isolated apps
  - can use `-p` or `--project-name` to specify a project name
- `docker compose up`: builds, (re)creates, starts, and attaches to containers for a service
- `docker compose down`: stops and removes containers, networks, volumes, and images created by `up`
- `docker compose ps`: lists containers
- `docker compose logs`: views output from containers
- `docker compose exec`: runs a command in a service
- `docker compose build`: builds or rebuilds services
- `docker compose pull`: pulls images for services
- `docker compose start`: starts existing containers for a service
- `docker compose stop`: stops existing containers without removing them
- `docker compose restart`: restarts services
- `docker compose pause`: pauses services
- `docker compose unpause`: unpauses services
