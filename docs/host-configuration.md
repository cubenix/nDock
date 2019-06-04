# Docker Host Configuration

## Prerequisites

- Docker is installed on the remote hosts
- Host can listen for TCP requests

## Setup 

The first thing you need, is to make the Docker daemon of your remote server listen for TCP requests at some port. Here we are going to make the daemon listen at port `2375`. 

I'm assuming that your remote hosts are running Unix based system. Here is what you need to do. Open the `docker service` file in an editor and update the value of `ExecStart`:

```
$ sudo nano /lib/systemd/system/docker.service

<snip>

ExecStart=/usr/bin/dockerd -H fd:// -H=tcp://0.0.0.0:2375

<snip>
```

Save the file and run the following commands to restart the Docker daemon:
```
$ sudo systemctl daemon-reload
$ sudo service docker restart
```

Now open a browser and try making a request to the URL: http://<host-ip>:2375/v1.38/containers/json. You should get a `json` response containing all your containers. If so, your host is all set.


**NOTE**: I also tried to make `Docker for Windows` listen to TCP requests. However, it didn't work out. Here is a [GitHub](https://github.com/docker/for-win/issues/225) link that has some good details. 
