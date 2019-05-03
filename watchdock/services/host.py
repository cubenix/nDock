""" Helper to work with a Docker host. """
import constants
import services.api as api


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
        client = api.create_client(host)
        host_containers[host.name] = len(client.containers.list())
    client.close()
    return host_containers


def get_containers_and_color_codes(host):
    """
    Returns
    [
        {
            "id": "container-id",
            "name": "container-name",
            "color_code": "color-code-for-container"
        }
    ]
    """
    client = api.create_client(host)
    container_list = client.containers.list()
    containers = []
    for index in range(len(container_list)):
        containers.append({
            "id": container_list[index].id,
            "name": container_list[index].name,
            "color_code": constants.COLOR_BACKGROUND[index]
        })
    client.close()
    return containers
