""" Helper to work with a Docker host. """
import docker
import json
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


def get_host_containers(host):
    """
    Returns
    [
        {
            "name": "container-name",
            "color_code": "color-code-for-container"
        }
    ]
    """
    client = __create_client(host)
    container_list = client.containers.list()
    containers = []
    for index in range(len(container_list)):
        containers.append({
            "name": container_list[index].name,
            "color_code": constants.COLOR_BACKGROUND[index]
        })
    return containers


def get_stats(host):
    client = __create_low_level_client(host)
    results = []
    for container in client.containers():
        stats = client.stats(container['Id'], stream=False)
        result = {}
        result["container"] = container["Names"][0].strip('/')
        result["usage"] = "{0:.3f}".format(__calculate_usage(stats))
        results.append(result)
    return json.dumps(results)


def __calculate_usage(stats):
    # get required readings
    cpu_stats = stats["cpu_stats"]["cpu_usage"]["total_usage"]
    precpu_stats = stats["precpu_stats"]["cpu_usage"]["total_usage"]
    system_cpu_stats = stats["cpu_stats"]["system_cpu_usage"]
    system_precpu_stats = stats["precpu_stats"]["system_cpu_usage"]
    cpu_count = len(stats["precpu_stats"]["cpu_usage"]["percpu_usage"])

    # calculate the change for cpu usage of the container in between readings
    cpu_delta = cpu_stats - precpu_stats

    # calculate the change for the entire system between readings
    system_delta = system_cpu_stats - system_precpu_stats

    cpu_usage = 0.0
    if system_delta > 0 and cpu_delta > 0:
        cpu_usage = (cpu_delta / system_delta) * cpu_count * 100
    return cpu_usage


def __create_client(host):
    # return docker.from_env()
    return docker.DockerClient(f"tcp://{host.IP}:{constants.DEFAULT_PORT}")


def __create_low_level_client(host):
    # return docker.APIClient()
    return docker.APIClient(f"tcp://{host.IP}:{constants.DEFAULT_PORT}")
