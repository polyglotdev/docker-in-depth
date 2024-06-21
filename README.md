# Docker in Depth

[Course Link](https://dominiquehallan-links.com/3zdb2MD)

## What is Docker?

Docker is a platform for developers and sysadmins to develop, deploy, and run applications with containers. The use of Linux containers to deploy applications is called containerization. Containers are not new, but their use for easily deploying applications is.

Containerization is increasingly popular because containers are:

- Flexible: Even the most complex applications can be containerized.
- Lightweight: Containers leverage and share the host kernel.
- Interchangeable: You can deploy updates and upgrades on-the-fly.
- Portable: You can build locally, deploy to the cloud, and run anywhere.
- Scalable: You can increase and automatically distribute container replicas.
- Stackable: You can stack services vertically and on-the-fly.

## What is a Container?

A container image is a lightweight, stand-alone, executable package of a piece of software that includes everything needed to run it: code, runtime, system tools, system libraries, settings. Available for both Linux and Windows-based applications, containerized software will always run the same, regardless of the environment. Containers isolate software from its surroundings, for example differences between development and staging environments and help reduce conflicts between teams running different software on the same infrastructure.

## What is an image?

An image is a lightweight, stand-alone, executable package that includes everything needed to run a piece of software, including the code, a runtime, libraries, environment variables, and config files.

## What is a `Dockerfile`?

A `Dockerfile` is a text document that contains all the commands a user could call on the command line to assemble an image. Using `docker build` users can create an automated build that executes several command-line instructions in succession.

## `docker commit`

You can turn an container into an image by using the `docker commit` command. This is not the best practice, but it can be useful for debugging or testing purposes. The `docker commit` command takes a container ID and creates an image from it.

You can also change instructions using the `--change` flag.

```bash
docker commit --change='CMD ["apache2ctl", "-D", "FOREGROUND"]' <container_id> <new_image_name>
```

## Docker image vs Docker container

A Docker image is a file, comprised of multiple layers, used to execute code in a Docker container. A Docker container is a runnable instance of a Docker image. You can create, start, stop, move, or delete a container using the Docker API or CLI. You can connect a container to one or more networks, attach storage to it, or even create a new image based on its current state.

## Port Mapping

Port mapping is a mechanism used by Docker to map a container's port to a port on the host machine. This allows external clients to access the containerized application by connecting to the host machine on the mapped port.

To map a port, you can use the `-p` flag when running a container.

```bash
docker run -p 80:80 <image_name>
```
