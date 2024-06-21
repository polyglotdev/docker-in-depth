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

You can also use:

```bash
docker run -P <image_name>
```

This will map all exposed ports in the container to random ports on the host machine.

### Docker Hierarchy and Lifecycle

1. **Dockerfile**:
   - A `Dockerfile` is a text file that contains a series of instructions on how to build a Docker image.
   - You write a `Dockerfile` to specify the base image, the software you want to install, and the configuration you need.

2. **Docker Image**:
   - An image is a read-only template that contains the application code, runtime, libraries, environment variables, and other dependencies needed to run an application.
   - You create an image by running the `docker build` command, which reads the instructions from the `Dockerfile`.
   - Images are stored in a registry (like Docker Hub) and can be pulled to different environments.

3. **Docker Container**:
   - A container is a runnable instance of an image. It is an isolated environment that runs the application defined by the image.
   - You create and start a container by running the `docker run` command on an image.
   - Containers are ephemeral by nature, meaning they can be started, stopped, moved, and deleted easily without affecting the underlying image.

### Lifecycle and Workflow

1. **Writing a Dockerfile**:
   - Begin by writing a `Dockerfile` with the necessary instructions to set up your application environment.
   - Example:
     ```dockerfile
     FROM golang:1.22.4-bookworm
     WORKDIR /app
     COPY . .
     RUN go build -o main .
     CMD ["./main"]
     ```

2. **Building an Image**:
   - Use the `docker build` command to create an image from your `Dockerfile`.
   - Example:
     ```bash
     docker build -t my-go-app .
     ```
   - This command reads the `Dockerfile` in the current directory (`.`) and builds an image tagged as `my-go-app`.

3. **Running a Container**:
   - Use the `docker run` command to create and start a container from the image.
   - Example:
     ```bash
     docker run -p 8080:8080 my-go-app
     ```
   - This command creates a container from the `my-go-app` image, maps port 8080 of the container to port 8080 on the host, and starts the application.

### Hierarchy Recap

1. **Dockerfile**: The starting point, a script with instructions to build an image.
2. **Docker Image**: Built from a Dockerfile, a read-only template for creating containers.
3. **Docker Container**: A running instance of an image, an isolated environment for running applications.

### Summary

- **First**, you write a `Dockerfile` with the necessary instructions.
- **Second**, you build a Docker image from the `Dockerfile`.
- **Third**, you run a container from the Docker image.

## Docker Networking Basics

Docker networking is a powerful feature that allows you to create virtual networks for your containers. By default, Docker creates a bridge network for you when you install it. This bridge network allows containers to communicate with each other and with the host machine.

> 💡 When using host networking whatever you open on the container is bound to the host.

### Docker Network Types

Docker provides several types of network drivers to support different networking use cases. Here's an overview of the main network types in Docker:

1. **Bridge Network**:
   - **Default bridge network**: Automatically created when Docker is installed. Containers on this network can communicate with each other using their IP addresses.
   - **User-defined bridge networks**: More flexible than the default bridge network, allowing containers to communicate by name rather than IP address. This is useful for creating isolated environments where multiple containers can interact.

2. **Host Network**:
   - In this mode, the container shares the host's networking namespace. There is no network isolation between the container and the host. This can be useful for performance-sensitive applications that require low latency, but it reduces security isolation.

3. **Overlay Network**:
   - Designed for multi-host networking, it allows containers on different Docker hosts to communicate with each other. This is particularly useful in swarm mode or other orchestrated environments where services need to span multiple hosts.

4. **None Network**:
   - This disables networking for a container. The container has no access to the network interfaces of the host and is isolated from other containers and the external network. This is useful for running tasks that do not require network access.

5. **MACVLAN Network**:
   - Assigns a MAC address to each container, making them appear as physical devices on the network. Containers can then be directly connected to the physical network. This is useful for legacy applications that require direct layer 2 access or for situations where network performance is critical.

6. **IPVLAN Network**:
   - Similar to MACVLAN but allows for more flexibility in IP address management and offers better performance for certain use cases. It can operate in different modes (l2, l3) depending on the desired network topology.

### Choosing the Right Network Type

- **Bridge**: Best for single-host deployments where you need simple communication between containers.
- **Host**: Use when performance and simplicity are more critical than security and isolation.
- **Overlay**: Ideal for multi-host deployments using Docker Swarm or other orchestration tools.
- **None**: Use when you need to completely isolate the container from any network.
- **Macvlan/IPvlan**: Suitable for complex network setups requiring direct access to the physical network or when higher network performance is needed.

### Example Commands

- **Creating a User-defined Bridge Network**:
  ```bash
  docker network create my_bridge_network
  ```

- **Running a Container on the Host Network**:
  ```bash
  docker run --network host my_image
  ```

- **Creating an Overlay Network**:
  ```bash
  docker network create --driver overlay my_overlay_network
  ```

- **Running a Container with No Network**:
  ```bash
  docker run --network none my_image
  ```

- **Creating a Macvlan Network**:
  ```bash
  docker network create -d macvlan \
    --subnet=192.168.1.0/24 \
    --gateway=192.168.1.1 \
    -o parent=eth0 macvlan_network
  ```

## Docker Volumes

Docker volumes are a way to persist data generated by and used by Docker containers. Volumes are stored outside the container's filesystem and are not deleted when the container is removed. This makes volumes a useful tool for managing data in Docker containers.

### Storage Options in Docker

1. Bind Mounts
2. Volumes
3. In memory storage: `tmpfs`

### Good use cases for bind mounts

In general, you should use volumes where possible. Bind mounts are appropriate for the following types of use case:

- Sharing configuration files from the host machine to containers. This is how Docker provides DNS resolution to containers by default, by mounting `/etc/resolv`.conf from the host machine into each container.
- Sharing source code or build artifacts bet a development environment on the Docker host and a container. For instance, you may mount a Maven `target/` directory into a container, and each time you build the Maven project on the Docker host, the container gets access to the rebuilt artifacts.
- If you use Docker for development this way, your production Dockerfile would copy the production-ready artifacts directly into the image, rather than relying on a bind mount.
- When the file or directory structure of the Docker host is guaranteed to be consistent with the bind mounts the containers require.
- When you need to be able to back up, restore, or migrate data from one Docker host to another, volumes are a better choice. You can stop containers using the volume, then back up the volume's directory (such as `/var/lib/docker/volumes/<volume-name>`).

### Good use cases for `tmpfs` mounts

`tmpfs` mounts are best used for cases when you do not want the data to persist either on the host machine or within the container. This be for security reasons or to protect the performance of the container when your application needs to write a large volume of non-persistent state data.

## Breakdown

```bash
➜ dkr -d --mount type=bind,src="./var/demo/logs",dst=/logs scratch_volume:latest
906aba69570bee3dc7f47151504dd903b27999e446e85f221dc307fa43480051
```

Here we are running a container with the `--mount` flag. We are specifying the type of mount as `bind`, the source directory as `./var/demo/logs`, and the destination directory as `/logs`. The image we are using is `scratch_volume:latest`.

The output of that file where the logs went was:

```log
Host: 906aba69570b
Loop Iteration: 0
Host: 906aba69570b
Loop Iteration: 1
Host: 906aba69570b
Loop Iteration: 2
Host: 906aba69570b
Loop Iteration: 3
Host: 906aba69570b
Loop Iteration: 4
Host: 906aba69570b
Loop Iteration: 5
Host: 906aba69570b
Loop Iteration: 6
Host: 906aba69570b
Loop Iteration: 7
Host: 906aba69570b
Loop Iteration: 8
Host: 906aba69570b
Loop Iteration: 9
Host: 906aba69570b
Loop Iteration: 10
Host: 906aba69570b
Loop Iteration: 11
Host: 906aba69570b
Loop Iteration: 12
Host: 906aba69570b
Loop Iteration: 13
Host: 906aba69570b
Loop Iteration: 14
Host: 906aba69570b
Loop Iteration: 15
Host: 906aba69570b
Loop Iteration: 16
Host: 906aba69570b
Loop Iteration: 17
Host: 906aba69570b
Loop Iteration: 18
Host: 906aba69570b
Loop Iteration: 19
```

So it worked as expected 🎆!

Let's look at a slightly different command with a different mount type:

```bash
docker-in-depth/docker/dockerfiles/writedata on  main [!] via 🐹 v1.22.4 on 🐳 v26.1.4 (orbstack) using ☁️  default/springboot-udemy-course
➜ dkr -d --mount type=volume,src="logs",dst=/logs scratch_volume:latest
a40c22f66f9917fc5dae9c2798a00c03de702530e034db1e7da4b88f9a10dfd5
```

Here we are running a container with the `--mount` flag. We are specifying the type of mount as `volume`, the source directory as `logs`, and the destination directory as `/logs`. The image we are using is `scratch_volume:latest`.
