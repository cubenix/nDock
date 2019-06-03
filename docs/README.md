# DockerDoodle

DockerDoodle started as a *fun side project*, hence the name, is an open-source project built around the idea of managing **Docker containers** running on **individual** hosts.
<br />

<br />Here is the first look:

![DockerDoodle](https://s3.gifyu.com/images/containers-count-min.gif)

## Features 

Below is a list of features in DockerDoodle upcoming releases:

### v0.1.0

- [x] Get the number of containers running on a host
  - [x] exclude stopped (`docker ps`)
  - [x] include stopped (`docker ps -a`)
- [x] Get stats from a Docker host (`docker stats`)
- [x] Get details of a container 
  - ID
  - Name
  - Image
  - Ports
  - Command
  - Mounts

### v0.2.0

- [ ] Add a Docker Registry 
- [ ] List images
- [ ] List image tags

### v0.3.0

- [ ] Perform containers operations
  - [ ] Start
  - [ ] Stop
  - [ ] Remove

### v0.4.0

- [ ] Perform Docker Image operations
  - [ ] Remove a tag
  - [ ] Add new tag
  - [ ] Delete an image

