import docker
import constants


def create_client(host):
    # return docker.from_env()
    return docker.DockerClient(f"tcp://{host.IP}:{constants.DEFAULT_PORT}")


def create_low_level_client(host):
    # return docker.APIClient()
    return docker.APIClient(f"tcp://{host.IP}:{constants.DEFAULT_PORT}")
