""" Helper to work with a Docker host. """
from docker import DockerClient
import constants


def get_containers(hosts):
    """Returns a dictionary with hostname and the number of containers in running state.

    :param hosts: list of Docker hosts defined in config.json

    Returns:
        { 'hostname': containers-count }
    """
    host_containers = {}
    for host in hosts:
        client = DockerClient(f"tcp://{host.IP}:{constants.DEFAULT_PORT}")
        host_containers[host.name] = len(client.containers.list())

    return host_containers
