""" Helper to work with a Docker host. """
import docker
from models.host import Host
import constants


def get_containers(hosts):
    """
    Returns a dictionary with hostname and the number of
    containers in running state.

    :param hosts: list of Docker hosts defined in config.json

    Returns:
        { 'hostname': containers-count }
    """
    host_containers = {}
    for host in hosts:
        client = __create_client(host)
        host_containers[host.name] = len(client.containers.list())
    return host_containers


def get_stats():
    host = Host(name="mglab-srv4", IP="172.27.127.134")
    client = __create_low_level_client(host)
    for container in client.containers():
        print(next(client.stats(resource_id=container['Id'], decode=True)))


def __create_client(host):
    return docker.DockerClient(f"tcp://{host.IP}:{constants.DEFAULT_PORT}")


def __create_low_level_client(host):
    return docker.api.client.APIClient(
        f"tcp://{host.IP}:{constants.DEFAULT_PORT}")
